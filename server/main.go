package main

import (
	"context"
	"github.com/BioChemML/SIC50/server/api"
	"github.com/BioChemML/SIC50/server/config"
	"github.com/BioChemML/SIC50/server/model"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()
	r := gin.New()
	err := api.Register(r)
	if err != nil {
		panic(err)
	}
	go func() {
		model.LoopTask(ctx)
	}()
	r.Run(config.Cfg.Addr)
}
