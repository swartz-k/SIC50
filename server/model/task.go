package model

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/BioChemML/SIC50/server/config"
	"github.com/BioChemML/SIC50/server/utils/files"
	"github.com/BioChemML/SIC50/server/utils/image"
	"github.com/BioChemML/SIC50/server/utils/log"
	"github.com/BioChemML/SIC50/server/utils/tensor"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"os"
	"path"
	"time"
)

type taskModel struct {
}

var TaskModel = taskModel{}

type TaskStatus string

const (
	TaskStatusPending TaskStatus = "PENDING"
	TaskStatusSuccess TaskStatus = "SUCCESS"
	TaskStatusFailed  TaskStatus = "FAILED"

	W int = 201
	H int = 201
)

type TaskInputUpload struct {
	Name string `json:"name"`
	Size string `json:"size"`
	Path string `json:"response"`
}

type TaskInput struct {
	Concentration float32            `json:"concentration"`
	Images        []*TaskInputUpload `json:"images"`
}

type TaskOutputResult struct {
	Con float32 `json:"con,omitempty"`
	Res float32 `json:"res,omitempty"`
}

type TaskOutput struct {
	Result []TaskOutputResult `json:"result"`
}

type TaskConfig struct {
	InputLayer  string             `json:"input_layer"`
	OutputLayer string             `json:"output_layer"`
	Steps       map[int]*TaskInput `json:"steps"`
	Output      *TaskOutput        `json:"output"`
}

func (t *TaskConfig) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal: %s", value)
	}

	return json.Unmarshal(bytes, t)
}

func (t TaskConfig) Value() (driver.Value, error) {
	return json.Marshal(t)
}

type Task struct {
	*gorm.Model
	TaskId string      `gorm:"type:varchar(100);column:task_id" json:"task_id"`
	Config *TaskConfig `json:"config"`
	Status TaskStatus  `json:"status"`
}

func (t *taskModel) getBaseModel() *gorm.DB {
	return db.Model(&Task{}).Preload(clause.Associations)
}

func (t *taskModel) GetByTaskId(taskId string) (*Task, error) {
	var task Task
	err := t.getBaseModel().Where("task_id = ?", taskId).First(&task).Error
	return &task, err
}

func (t *taskModel) GetByStatus(status TaskStatus) (*Task, error) {
	var task Task
	err := t.getBaseModel().Where("status = ?", status).First(&task).Error
	return &task, err
}

func (t *taskModel) Save(ctx context.Context, task *Task) error {
	return t.getBaseModel().WithContext(ctx).Save(task).Error
}

func (t *taskModel) Count(ctx context.Context) (*int64, error) {
	var r int64
	err := t.getBaseModel().WithContext(ctx).Count(&r).Error
	return &r, err
}

func (t *taskModel) Update(ctx context.Context, task *Task) error {
	return t.getBaseModel().WithContext(ctx).Where("task_id = ?", task.TaskId).Updates(task).Error
}

func (t *Task) MockCal(ctx context.Context) (*Task, error) {
	input := "serving_default_input_input"
	output := "StatefulPartitionedCall"
	image := "$GOPATH/github.com/swartz-k/Train_test_images/train/con1/edge_B16_Paclitaxel_ctrl_con1_20220704_C07_sx_1_sy_1_w1-001.png"
	var results []TaskOutputResult
	for i, inp := range t.Config.Steps {
		res, err := tensor.Cal(image, input, output)
		if err != nil {
			return nil, err
		}
		results = append(results, TaskOutputResult{Con: inp.Concentration, Res: res[0].Value().([][]float32)[0][1]})
		log.Info("step %d results %+v", i, results)
	}
	t.Config.Output.Result = results
	t.Status = TaskStatusSuccess
	err := TaskModel.getBaseModel().WithContext(ctx).Where("id = ?", t.ID).Save(t).Error
	return nil, err
}

func (t *Task) ReTrain(ctx context.Context) (*Task, error) {
	/*
		prepare task input data
		Steps[0] ctl compare
		Steps[1] compare with ctl
	*/
	basePath := path.Join(config.Cfg.UploadDir, t.TaskId)

	defer func() {
		_ = TaskModel.Update(ctx, t)
	}()
	// prepare step
	var minDirTotalImage int
	for k, s := range t.Config.Steps {
		// ctrl not need dir
		if k == 0 {
			continue
		}
		// step path
		basSPath := path.Join(basePath, fmt.Sprintf("%d", k))
		sPath := path.Join(basSPath, fmt.Sprintf("%d", k))
		err := os.MkdirAll(sPath, 0777)
		if err != nil {
			return nil, errors.Wrapf(err, "mkdir for task step %s", sPath)
		}
		// step path and compare path
		for _, img := range s.Images {
			tmpImg := img
			var tmpMin int
			_, fName := path.Split(tmpImg.Path)
			tPath := fmt.Sprintf("%s.png", path.Join(sPath, fName))
			log.Info("%d, mv img %s to %s", k, tmpImg, tPath)
			err = files.Copy(tmpImg.Path, tPath)
			if err != nil {
				return nil, errors.Wrapf(err, "prepare step copy %s => %s", tmpImg.Path, tPath)
			}
			log.Info("err %+v", err)
			tmpMin, err = image.Split(tPath, W, H)
			if err == nil {
				log.Info("should remove %s", tPath)
				_ = os.Remove(tPath)
				if tmpMin < minDirTotalImage {
					minDirTotalImage = tmpMin
				}
			}
		}
		// ctrl image
		ctrlPath := path.Join(basSPath, "0")
		err = os.MkdirAll(ctrlPath, 0777)
		if err != nil {
			return nil, errors.Wrapf(err, "mkdir for task ctrl step %s", ctrlPath)
		}
		for _, img := range t.Config.Steps[0].Images {
			tmpImg := img
			var tmpMin int
			_, fName := path.Split(tmpImg.Path)
			tFile := fmt.Sprintf("%s.png", path.Join(ctrlPath, fName))
			err = files.Copy(tmpImg.Path, tFile)
			if err != nil {
				return nil, errors.Wrapf(err, "prepare ctrl copy %s => %s", tmpImg.Path, tFile)
			}
			tmpMin, err = image.Split(tFile, W, H)
			if err == nil {
				_ = os.Remove(tFile)
				if tmpMin < minDirTotalImage {
					minDirTotalImage = tmpMin
				}
			}
		}
	}

	// 1. Train Script
	output := &TaskOutput{Result: nil}
	for k, s := range t.Config.Steps {
		if k == 0 {
			continue
		}
		trainPath := path.Join(basePath, fmt.Sprintf("%d", k))
		r, err := tensor.Train(trainPath, minDirTotalImage)

		if err != nil {
			return nil, errors.Wrap(err, "cal tensor")
		}
		output.Result = append(output.Result, TaskOutputResult{Con: s.Concentration, Res: *r})
	}
	t.Config.Output = output
	if len(output.Result)+1 == len(t.Config.Steps) {
		t.Status = TaskStatusSuccess
	}
	return t, nil
}

func (t *Task) Cal(ctx context.Context) (*Task, error) {
	/*
		prepare task input data
		Steps[0] ctl compare
		Steps[1] compare with ctl
	*/
	basePath := path.Join(config.Cfg.UploadDir, t.TaskId)
	if _, err := os.Open(basePath); os.IsNotExist(err) {
		_ = os.Mkdir(t.TaskId, 0644)
	}
	defer func() {
		_ = TaskModel.Update(ctx, t)
	}()
	for k, s := range t.Config.Steps {
		sPath := path.Join(basePath, fmt.Sprintf("%d", k))
		err := os.MkdirAll(sPath, 0777)
		if err != nil {
			return nil, errors.Wrapf(err, "mkdir for task step %s", sPath)
		}
		for _, img := range s.Images {
			tmpImg := img
			_, fName := path.Split(tmpImg.Path)
			tPath := path.Join(sPath, fName)
			err = os.Rename(tmpImg.Path, fmt.Sprintf("%s.png", tPath))
			if err != nil {
				return nil, errors.Wrapf(err, "move %s => %s", tmpImg.Path, tPath)
			}
			tmpImg.Path = tPath
		}
	}

	// 0. image to tensor && generate result
	output := &TaskOutput{Result: nil}
	for _, s := range t.Config.Steps {
		var result *TaskOutputResult
		for _, img := range s.Images {
			r, err := tensor.Call(img.Path)
			if err != nil {
				return nil, errors.Wrap(err, "cal tensor")
			}
			// fixme: ont step multi image
			f := r[0].Value().([][]float32)[0][1]
			result = &TaskOutputResult{Con: s.Concentration, Res: f}
		}
		output.Result = append(output.Result, *result)
	}
	t.Config.Output = output

	// 1. Train Script
	return t, nil
}

func LoopTask(ctx context.Context) {
	t := time.NewTicker(time.Second * 3)
	for {
		select {
		case <-t.C:
			task, err := TaskModel.GetByStatus(TaskStatusPending)
			if err != nil {
				log.Info("get pending task failed %+v", err)
				break
			}
			_, err = task.ReTrain(ctx)
			if err != nil {
				log.Info("cal task failed %+v", err)
				break
			}
		}
	}
}
