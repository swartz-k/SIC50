package image

import (
	"testing"
)

func TestSplit(t *testing.T) {
	img := "test/test.png"

	min, err := Split(img, 201, 201)
	t.Logf("img %s total %d, err %+v", img, min, err)
}

func Test_generateName(t *testing.T) {
	a := "tmp/a.b.c.png"
	b := "/a/b/c/da.png"
	c := "aaaaa"
	suffix := "1-1"
	t.Logf("name for %s if %s", a, generateName(a, suffix))
	t.Logf("name for %s if %s", b, generateName(b, suffix))
	t.Logf("name for %s if %s", c, generateName(c, suffix))
}
