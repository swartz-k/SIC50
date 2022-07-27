package image

import (
	"github.com/pkg/errors"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"io/ioutil"
)

func Transfer(p string) (*tf.Tensor, error) {
	_, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, errors.Wrapf(err, "transfer img %s", p)
	}

	//tf.ReadTensor(tf.Float, ) [1][198][198][3]float32{})
	return tf.NewTensor([1][198][198][3]float32{})
	//return tf.ReadTensor(tf.Float, []int64{198, 198, 3}, bytes.NewReader(c))
	// return tf.NewTensor(c)
}
