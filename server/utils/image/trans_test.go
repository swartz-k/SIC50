package image

import "testing"

func Test_Trans(t *testing.T) {
	f := "/Users/joker/work/python/tools/sic50_in/xb-yhll.png"
	err := Transfer(f)
	if err != nil {
		t.Fatal(err)
	}
}
