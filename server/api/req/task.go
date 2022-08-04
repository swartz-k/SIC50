package req

import (
	"github.com/BioChemML/SIC50/server/model"
	"github.com/google/uuid"
)

type CreateTaskContent struct {
	Concentration float32                  `json:"concentration"`
	Upload        []*model.TaskInputUpload `json:"upload"`
}

type CreateTask struct {
	Input  string               `json:"input,omitempty"`
	Output string               `json:"output,omitempty"`
	Step   []*CreateTaskContent `json:"step"`
}

func (r *CreateTask) GetTask() *model.Task {
	if r.Input == "" {
		r.Input = "serving_default_input_input"
	}
	if r.Output == "" {
		r.Output = "StatefulPartitionedCall"
	}
	steps := make(map[int]*model.TaskInput, len(r.Step))

	for k, i := range r.Step {
		c := &model.TaskInput{}
		c.Concentration = i.Concentration
		c.Images = i.Upload
		steps[k] = c
	}

	t := &model.Task{
		TaskId: uuid.New().String(),
		Config: &model.TaskConfig{
			InputLayer:  r.Input,
			OutputLayer: r.Output,
			Output:      &model.TaskOutput{},
			Steps:       steps,
		},
		Status: model.TaskStatusPending,
	}
	return t
}
