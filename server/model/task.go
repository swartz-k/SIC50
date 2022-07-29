package model

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/BioChemML/SIC50/server/config"
	"github.com/BioChemML/SIC50/server/utils/tensor"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"os"
	"path"
)

type taskModel struct {
}

var TaskModel = taskModel{}

type TaskStatus string

const (
	TaskStatusPending TaskStatus = "PENDING"
	TaskStatusSuccess TaskStatus = "SUCCESS"
	TaskStatusFailed  TaskStatus = "FAILED"
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

type TaskOutput struct {
	Result string `json:"result"`
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

func (t *taskModel) Save(ctx context.Context, task *Task) error {
	return t.getBaseModel().WithContext(ctx).Save(task).Error
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
		_ = TaskModel.Save(ctx, t)
	}()
	for k, s := range t.Config.Steps {
		sPath := path.Join(basePath, fmt.Sprintf("%d", k))
		err := os.MkdirAll(sPath, 0777)
		if err != nil {
			return nil, errors.Wrapf(err, "mkdir for task step %s", sPath)
		}
		for _, img := range s.Images {
			_, fName := path.Split(img.Path)
			tPath := path.Join(sPath, fName)
			err = os.Rename(img.Path, tPath)
			if err != nil {
				return nil, errors.Wrapf(err, "move %s => %s", img.Path, tPath)
			}
			img.Path = tPath
		}
	}

	r, err := tensor.Cal(basePath, t.Config.InputLayer, t.Config.OutputLayer)
	if err != nil {
		return nil, errors.Wrap(err, "cal tensor")
	}

	output := &TaskOutput{Result: fmt.Sprintf("%s", r[0].Value())}
	t.Config.Output = output
	err = TaskModel.Save(ctx, t)
	if err != nil {
		return nil, errors.Wrap(err, "save task")
	}
	return t, nil
}
