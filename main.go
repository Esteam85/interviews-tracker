package main

import (
	"github.com/esteam85/interviews-tracker/process/infrastructure/gin"
	"github.com/esteam85/interviews-tracker/process/infrastructure/log"
)

func main() {
	router := gin.NewEngine()
	processHandler := gin.ProcessHandler{}
	router.POST("/processes", processHandler.AddProcessHandler)
	err := router.Run("localhost:8080")
	if err != nil {
		log.Error("error trying to run server", err)
		panic("http server fail!")
	}
}
