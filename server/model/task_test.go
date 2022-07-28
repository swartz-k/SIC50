package model

import (
	"context"
	"testing"
)

func Test_TaskCal(t *testing.T) {
	ctx := context.Background()
	task, err := TaskModel.GetByTaskId("ae7c75db-12d1-4e4e-955c-405b250c783a")
	if err != nil {
		t.Fatal(err)
	}
	task, err = task.Cal(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", task)
}
