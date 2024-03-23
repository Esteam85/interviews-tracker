package main

import (
	"github.com/esteam85/interviews-tracker/process/infrastructure/gin"
)

func main() {
	router := gin.NewEngine()
	processHandler := gin.ProcessHandler{}
	router.POST("/processes", processHandler.AddProcessHandler)
	err := router.Run("localhost:8080")
	if err != nil {
		panic("http server fail!")
	}
}
