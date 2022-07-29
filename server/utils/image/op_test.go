package image

import (
	"math"
	"testing"
)

func Test_calPow(t *testing.T) {
	t.Logf("2 ** 10 = %f", math.Pow(2, 10))
	t.Logf("3 ** 0.8 = %f", math.Pow(3, 0.8))
}
