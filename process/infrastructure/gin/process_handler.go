package gin

import (
	"context"
	"errors"
	"net/http"

	"github.com/esteam85/interviews-tracker/process/infrastructure/log"

	"github.com/esteam85/interviews-tracker/process/service"

	"github.com/esteam85/interviews-tracker/process/domain"

	"github.com/gin-gonic/gin"
)

type ProcessHandler struct {
	service *service.ProcessService
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
		pAsPrimitives.Id,
		pAsPrimitives.PostulationType,
		pAsPrimitives.Platform,
		pAsPrimitives.Company,
		pAsPrimitives.Position,
		pAsPrimitives.JobType,
		options...)
	if err != nil {
		handleError(c, err)
	}
	c.JSON(http.StatusCreated, pAsPrimitives)
}

func handleError(c *gin.Context, err error) {

	switch {
	case errors.Is(err, domain.ErrInvalidProcessID):
		c.JSON(http.StatusInternalServerError, err.Error())
	default:
		log.Error("internal server error,", err.Error())
		c.JSON(http.StatusInternalServerError, err)
	}
}
