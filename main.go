package main

import (
	"github.com/esteam85/interviews-tracker/process/infrastructure"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/processes", infrastructure.AddProcessHandler)
	err := router.Run("localhost:8080")
	if err != nil {
		panic("http server fail!")
	}
}
