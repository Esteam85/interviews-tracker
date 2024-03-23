package log

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	logger = zap.Must(zap.NewProduction())
	if os.Getenv("APP_ENV") == "development" {
		logger = zap.Must(zap.NewDevelopment())
	}
}

func Info(msg string, err error) {
	logger.Info(fmt.Sprintf("%s: %s", msg, err.Error()))
}

func Error(msg string, err error) {
	logger.Error(fmt.Sprintf("%s: %s", msg, err.Error()))
}
