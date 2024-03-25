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
	var options []domain.ProcessOptions
	if pAsPrimitives.FirstContact != nil {
		options = append(options, domain.WithFirstContact(pAsPrimitives.FirstContact))
	}

	if pAsPrimitives.Salary != nil {
		options = append(options, domain.WithSalary(pAsPrimitives.Salary))
	}

	if pAsPrimitives.Client != "" {
		options = append(options, domain.WithClient(pAsPrimitives.Client))
	}

	err = p.service.AddProcess(ctx,
		pAsPrimitives.ProcessID,
		pAsPrimitives.PostulationType,
		pAsPrimitives.Platform,
		pAsPrimitives.Company,
		pAsPrimitives.Position,
		pAsPrimitives.JobType,
		options...)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, pAsPrimitives)
}

func handleError(c *gin.Context, err error) {

	switch {
	case errors.Is(err, domain.ErrInvalidCurrency),
		errors.Is(err, domain.ErrInvalidProcessID):
		c.String(http.StatusInternalServerError, err.Error())
	default:
		log.Error("internal server error,", err.Error())
		c.String(http.StatusInternalServerError, err.Error())
	}
}
