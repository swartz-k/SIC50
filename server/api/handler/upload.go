package handler

import (
	"github.com/BioChemML/SIC50/server/model"
	"github.com/BioChemML/SIC50/server/utils/log"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"path"
)

var dataPath = "data"

type ContentUpload struct {
	Response string `json:"response"`
}

type Content struct {
	Num    float32         `json:"num"`
	Upload []ContentUpload `json:"upload"`
}

type Req struct {
	Input   string          `json:"input,omitempty"`
	Output  string          `json:"output,omitempty"`
	Content map[int]Content `json:"content"`
}

func Cal(c *gin.Context) {
	req := Req{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	if req.Input == "" {
		req.Input = "serving_default_input_input"
	}
	if req.Output == "" {
		req.Output = "StatefulPartitionedCall"
	}
	log.Info("cal json %+v", req.Content)
	//logrus.
	var result []interface{}
	for _, con := range req.Content {
		for _, p := range con.Upload {
			tensors, err := model.GenerateTensorResult(req.Input, req.Output, p.Response)
			if err != nil {
				log.Info("failed with %s %+v", p.Response, err)
				c.JSON(http.StatusInternalServerError, err.Error())
				return
			}
			result = append(result, tensors[0].Value())
		}
	}
	c.JSON(http.StatusOK, result)
}

func Upload(c *gin.Context) {
	h, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	s, _ := c.GetPostForm("image")

	log.Info("map %+v, string %s, array %+v", c.PostFormMap("image"), s, c.PostFormArray("image"))
	f, err := h.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	content, err := ioutil.ReadAll(f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	fPath := path.Join(dataPath, uuid.New().String())
	err = ioutil.WriteFile(fPath, content, 0644)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, fPath)
}
