package api

import (
	"github.com/BioChemML/SIC50/server/api/handler"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func Register(r *gin.Engine) error {
	_, err := GetAuthMiddleware()
	if err != nil {
		return errors.Wrap(err, "get auth middle ware")
	}

	r.GET("/api/v1/overview", handler.Overview)
	r.POST("/api/v1/upload", handler.Upload)

	r.GET("/api/v1/task", handler.GetTask)
	r.POST("/api/v1/task", handler.Cal)
	r.POST("/api/v1/task/async", handler.AsyncCal)
	return nil
}
