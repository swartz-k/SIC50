package model

import (
	"fmt"
	"github.com/BioChemML/SIC50/server/utils/image"
	"github.com/BioChemML/SIC50/server/utils/log"
	"github.com/pkg/errors"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

var model *tf.SavedModel

func init() {
	var err error
	model, err = tf.LoadSavedModel("../model_p", []string{"serve"}, nil)
	if err != nil {
		panic(err)
	}
	defer model.Session.Close()
}

func GenerateTensorResult(input, output string, imgPath string) ([]*tf.Tensor, error) {
	model, err := tf.LoadSavedModel("../model_p", []string{"serve"}, nil)
	if err != nil {
		return nil, err
	}
	defer model.Session.Close()
	for _, i := range model.Graph.Operations() {
		fmt.Printf("operation name: %s, inputs: %d, outputs: %d, type: %s \n", i.Name(), i.NumInputs(), i.NumOutputs(), i.Type())
	}
	tensor, err := image.Transfer(imgPath)
	//tensor, err := tf.NewTensor([1][198][198][3]float32{})
	if err != nil {
		return nil, errors.Wrap(err, "deal with image tra")
	}
	log.Info("deal with image %s\n", imgPath)
	i := model.Graph.Operation(input)
	if i != nil {
		fmt.Printf("iiii %s input %d, output %d \n", i.Name(), i.NumInputs(), i.NumOutputs())
	} else {
		fmt.Printf("input op %s nil\n", input)
	}
	o := model.Graph.Operation(output)
	if o != nil {
		fmt.Printf("oooo %s input %d, output %d \n", o.Name(), o.NumInputs(), o.NumOutputs())
	} else {
		fmt.Printf("output op %s nil\n", output)
	}

	result, err := model.Session.Run(
		map[tf.Output]*tf.Tensor{
			//model.Graph.Operation("conv2d/kernel").Output(0): tensor, // Replace this with your input layer name
			model.Graph.Operation(input).Output(0): tensor, // Replace this with your input layer name
		},
		[]tf.Output{
			model.Graph.Operation(output).Output(0), // Replace this with your output layer name
		},
		nil,
	)
	return result, err
}
