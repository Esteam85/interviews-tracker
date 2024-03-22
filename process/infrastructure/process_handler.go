package infrastructure

import (
	"github.com/esteam85/interviews-tracker/process/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddProcessHandler(c *gin.Context) {
	var process *domain.Process
	err := c.BindJSON(&process)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, c.Error(err))
	}
}
