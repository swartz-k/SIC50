package model

import (
	"context"
	"testing"
)

func Test_TaskCal(t *testing.T) {
	ctx := context.Background()
	task, err := TaskModel.GetByTaskId("0417ff14-a686-4369-a588-7de9c96861f9")
	if err != nil {
		t.Fatal(err)
	}
	task, err = task.Cal(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", task)
}
