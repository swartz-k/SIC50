package tensor

import (
	"os"
	"testing"
)

func Test_getEpochStep(t *testing.T) {
	epoch, step := getEpochStep("", 10)
	t.Logf("r for %d epoch: %d, step:%d", 10, epoch, step)

	epoch, step = getEpochStep("", 33)
	t.Logf("r for %d epoch: %d, step:%d", 33, epoch, step)

	epoch, step = getEpochStep("", 3)
	t.Logf("r for %d epoch: %d, step:%d", 3, epoch, step)
}

func Test_Train(t *testing.T) {
	p := os.Getenv("TEST_PATH")
	re, err := Train(p, 3)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("result %f", *re)
}
