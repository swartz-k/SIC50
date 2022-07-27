package main

import (
	"github.com/BioChemML/SIC50/server/api"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.New()
	api.Register(r)
	r.Run()
}
