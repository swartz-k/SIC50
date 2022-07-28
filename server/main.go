package main

import (
	"github.com/BioChemML/SIC50/server/api"
	"github.com/BioChemML/SIC50/server/config"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	err := api.Register(r)
	if err != nil {
		panic(err)
	}
	r.Run(config.Cfg.Addr)
}
