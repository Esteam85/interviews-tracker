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

func Info(i ...interface{}) {
	logger.Info(fmt.Sprint(i...))
}

func Error(i ...interface{}) {
	logger.Error(fmt.Sprint(i...))
}

func Errorf(format string, i ...interface{}) {
	logger.Error(fmt.Sprintf(format, i...))
}

func Fatal(i ...interface{}) {
	logger.Fatal(fmt.Sprint(i...))
}
