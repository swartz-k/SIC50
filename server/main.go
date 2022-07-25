package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"net/http"
)

func main() {
	r := gin.New()

	r.POST("/api/v1/upload", HandleInput)
	r.Run("127.0.0.1:8081")
}

type Req struct {
	Input  string `json:"input,omitempty"`
	Output string `json:"output"`
}

func HandleInput(c *gin.Context) {
	req := Req{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	model, err := tf.LoadSavedModel("../model", []string{"serve"}, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer model.Session.Close()

	for _, i := range model.Graph.Operations() {
		fmt.Printf("operation name: %s, inputs: %d, outputs: %d, type: %s \n", i.Name(), i.NumInputs(), i.NumOutputs(), i.Type())
	}
	tensor, err := tf.NewTensor([1][198][198][3]float32{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	i := model.Graph.Operation("serving_default_conv2d_input")
	fmt.Printf("%s input %d, output %d \n", i.Name(), i.NumInputs(), i.NumOutputs())

	result, err := model.Session.Run(
		map[tf.Output]*tf.Tensor{
			//model.Graph.Operation("conv2d/kernel").Output(0): tensor, // Replace this with your input layer name
			model.Graph.Operation(req.Input).Output(0): tensor, // Replace this with your input layer name
		},
		[]tf.Output{
			model.Graph.Operation(req.Output).Output(0), // Replace this with your output layer name
		},
		nil,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			fmt.Sprintf("Error running the session with input, err: %s\n", err.Error()))
	}
	c.JSON(http.StatusOK, fmt.Sprintf("%+v", result[0].Value()))
}
