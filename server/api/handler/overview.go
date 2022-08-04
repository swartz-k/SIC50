package handler

import (
	"github.com/BioChemML/SIC50/server/config"
	"github.com/BioChemML/SIC50/server/model"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type OverviewResponse struct {
	Tasks  int64 `json:"tasks,omitempty"`
	Images int   `json:"images,omitempty"`
}

func Overview(c *gin.Context) {
	fs, err := ioutil.ReadDir(config.Cfg.UploadDir)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	total, err := model.TaskModel.Count(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	r := OverviewResponse{
		Tasks:  *total,
		Images: len(fs),
	}
	c.JSON(http.StatusOK, r)
}
