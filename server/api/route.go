package api

import (
	"github.com/BioChemML/SIC50/server/api/handler"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.POST("/api/v1/upload", handler.Upload)
	r.POST("/api/v1/handle", handler.Cal)
}
