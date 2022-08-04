package handler

import (
	"fmt"
	"github.com/BioChemML/SIC50/server/api/req"
	"github.com/BioChemML/SIC50/server/model"
	"github.com/BioChemML/SIC50/server/utils/log"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

func GetTask(c *gin.Context) {
	taskId := c.Query("id")
	// sql
	checkLength := 36
	if len(taskId) != checkLength {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("task id should be %d length", checkLength))
		return
	}
	t, err := model.TaskModel.GetByTaskId(taskId)
	log.Info("get task by id %s, %+v", taskId, t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.Wrap(err, "get task"))
		return
	}
	c.JSON(http.StatusOK, t)
}

func AsyncCal(c *gin.Context) {
	r := req.CreateTask{}
	err := c.BindJSON(&r)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	t := r.GetTask()
	err = model.TaskModel.Save(c, t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	//var result []interface{}
	//for _, con := range req.Content {
	//	for _, p := range con.Upload {
	//		tensors, err := model.GenerateTensorResult(req.Input, req.Output, p.Response)
	//		if err != nil {
	//			log.Info("failed with %s %+v", p.Response, err)
	//			c.JSON(http.StatusInternalServerError, err.Error())
	//			return
	//		}
	//		result = append(result, tensors[0].Value())
	//	}
	//}
	c.JSON(http.StatusOK, t)
}

func Cal(c *gin.Context) {
	req := req.CreateTask{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	t := req.GetTask()
	err = model.TaskModel.Save(c, t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	t, err = t.Cal(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, t.Config.Output.Result)
}
