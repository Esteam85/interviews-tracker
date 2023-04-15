package service_test

import (
	"errors"
	"github.com/esteam85/interviews-tracker/process/domain"
	"github.com/esteam85/interviews-tracker/process/domain/mocks"
	"github.com/esteam85/interviews-tracker/process/service"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddProcessSuccessfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := mock_domain.NewMockRepository(ctrl)
	m.EXPECT().Save(gomock.Any()).Return(nil)
	processUsecase := service.NewProcessService(m)
	err := processUsecase.AddProcess(uuid.New().String(),
		"Own",
		"Linkedin",
		"Esteam",
		"Dev",
		"contract",
		domain.WithSalary(6000, "usd"))
	assert.NoError(t, err)
}

func TestAddProcessWithError(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := mock_domain.NewMockRepository(ctrl)
	m.EXPECT().Save(gomock.Any()).Return(errors.New("error"))
	processUsecase := service.NewProcessService(m)
	err := processUsecase.AddProcess(uuid.New().String(),
		"Own",
		"Linkedin",
		"Esteam",
		"Dev",
		"contract",
		domain.WithSalary(6000, "usd"))
	assert.EqualError(t, err, "error")
}
