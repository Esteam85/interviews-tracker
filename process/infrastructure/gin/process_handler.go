package gin

import (
	"net/http"

	"github.com/esteam85/interviews-tracker/process/domain"

	"github.com/gin-gonic/gin"
)

type ProcessHandler struct {
}

func (p *ProcessHandler) AddProcessHandler(c *gin.Context) {
	var process *domain.Process
	err := c.BindJSON(&process)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, c.Error(err))
	}
}
