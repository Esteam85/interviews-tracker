package main

import (
	"github.com/esteam85/interviews-tracker/process/infrastructure/gin"
	"github.com/esteam85/interviews-tracker/process/infrastructure/log"
	"github.com/esteam85/interviews-tracker/process/infrastructure/mongo"
	"github.com/esteam85/interviews-tracker/process/service"
)

func main() {
	router := gin.NewEngine()
	mongoClient, err := mongo.NewClient()
	if err != nil {
		panic("http server fail!")
	}
	defer mongoClient.Disconnect()

	processMongoRepository := mongo.NewProcessMongoRepository(mongoClient.Client())
	processService := service.NewProcessService(processMongoRepository)
	processHandler := gin.NewProcessService(processService)
	router.POST("/processes", processHandler.AddProcessHandler)
	router.GET("/processes", processHandler.GetAllProcesses)

	log.Info("starting server...")
	err = router.Run("0.0.0.0:8080")
	if err != nil {
		log.Error("error trying to run server", err)
		panic("http server fail!")
	}
}
