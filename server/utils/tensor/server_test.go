package tensor

import (
	"io/ioutil"
	"path"
	"testing"
)

func Test_Cal(t *testing.T) {
	input := "serving_default_input_input"
	output := "StatefulPartitionedCall"
	image := "/Users/reachy/go/src/github.com/swartz-k/Train_test_images/train/con1/edge_B16_Paclitaxel_ctrl_con1_20220704_C07_sx_1_sy_1_w1-001.png"
	r, err := Cal(image, input, output)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("r %+v", r)
	t.Logf("result %+v", r[0].Value())
}

func Test_batchCall(t *testing.T) {
	//image := "/Users/reachy/go/src/github.com/swartz-k/Train_test_images/train/con1/edge_B16_Paclitaxel_ctrl_con1_20220704_C07_sx_1_sy_1_w1-001.png"
	p := "/Users/reachy/go/src/github.com/swartz-k/Train_test_images/train/con1/"
	fs, err := ioutil.ReadDir(p)
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range fs {
		a := path.Join(p, f.Name())
		r, err := Call(a)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("result %s %+v", a, r[0].Value().([][]float32)[0][1])

	}
}

func Test_call(t *testing.T) {
	//image := "/Users/reachy/go/src/github.com/swartz-k/Train_test_images/train/con1/edge_B16_Paclitaxel_ctrl_con1_20220704_C07_sx_1_sy_1_w1-001.png"
	p := "/Users/reachy/go/src/github.com/swartz-k/Train_test_images/train/con1/edge_B16_Paclitaxel_ctrl_con1_20220704_C07_sx_1_sy_1_w1-001.png"

	r, err := Call(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("result %s %+v", p, r[0].Value().([][]float32)[0][1])

}
