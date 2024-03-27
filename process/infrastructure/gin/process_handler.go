package gin

import (
	"context"
	"errors"
	"net/http"

	"github.com/esteam85/interviews-tracker/process/domain"
	"github.com/esteam85/interviews-tracker/process/infrastructure/log"
	"github.com/esteam85/interviews-tracker/process/service"

	"github.com/gin-gonic/gin"
)

type ProcessHandler struct {
	service *service.ProcessService
}

func NewProcessService(service *service.ProcessService) *ProcessHandler {
	return &ProcessHandler{
		service: service,
	}
}

func (p *ProcessHandler) AddProcessHandler(c *gin.Context) {
	ctx := context.Background()
	var pAsPrimitives *domain.ProcessAsPrimitives
	err := c.BindJSON(&pAsPrimitives)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, c.Error(err))
	}
	err = p.service.AddProcess(ctx,
		pAsPrimitives.ProcessID,
		pAsPrimitives.PostulationType,
		pAsPrimitives.Platform,
		pAsPrimitives.Company,
		pAsPrimitives.Position,
		pAsPrimitives.JobType,
		domain.WithFirstContact(pAsPrimitives.FirstContact),
		domain.WithSalary(pAsPrimitives.Salary),
		domain.WithClient(pAsPrimitives.Client))
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": pAsPrimitives})
}

func (p *ProcessHandler) GetAllProcesses(c *gin.Context) {
	processes, err := p.service.GetAllProcesses(context.Background())
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": processes.ToProcessesAsPrimitives()})
}

func handleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, domain.ErrInvalidCurrency),
		errors.Is(err, domain.ErrInvalidProcessID):
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	case errors.Is(err, domain.ErrProcessAlreadyExist):
		c.IndentedJSON(http.StatusConflict, gin.H{"message": err.Error()})
	default:
		log.Error("internal server error,", err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
	}
}
