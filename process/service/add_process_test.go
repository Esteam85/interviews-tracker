package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/esteam85/interviews-tracker/process/domain"
	mockdomain "github.com/esteam85/interviews-tracker/process/domain/mocks"
	"github.com/esteam85/interviews-tracker/process/service"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAddProcessSuccessfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := mockdomain.NewMockRepository(ctrl)
	m.EXPECT().Save(context.Background(), gomock.Any()).Return(nil)
	processUsecase := service.NewProcessService(m)
	err := processUsecase.AddProcess(
		context.TODO(),
		uuid.New().String(),
		"Own",
		"Linkedin",
		"Esteam",
		"Dev",
		"contract",
		domain.WithSalary(&domain.SalaryAsPrimitives{Amount: 6000, Currency: "usd", SalaryType: "net", SalaryPeriod: "yearly"}))
	assert.NoError(t, err)
}

func TestAddProcessWithError(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := mockdomain.NewMockRepository(ctrl)
	m.EXPECT().Save(context.Background(), gomock.Any()).Return(errors.New("error"))
	processUsecase := service.NewProcessService(m)
	err := processUsecase.AddProcess(context.Background(), uuid.New().String(),
		"Own",
		"Linkedin",
		"Esteam",
		"Dev",
		"contract",
		domain.WithSalary(&domain.SalaryAsPrimitives{Amount: 6000, Currency: "usd", SalaryType: "gross", SalaryPeriod: "monthly"}))
	assert.EqualError(t, err, "error")
}
