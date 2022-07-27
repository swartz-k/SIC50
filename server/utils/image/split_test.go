package image

import (
	"fmt"
	"os"
	"testing"
)

func TestSplit(t *testing.T) {
	img, err := os.Open("edge_B16_Paclitaxel_ctrl_con1_20220703-1_C07_sx_1_sy_1_w1-043.png")
	if err != nil {
		t.Fatal(err)
	}
	defer img.Close()
	fmt.Println(Split(img))
}
