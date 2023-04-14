package service_test

import (
	"github.com/esteam85/interviews-tracker/process/domain"
	"github.com/esteam85/interviews-tracker/process/domain/mocks"
	"github.com/esteam85/interviews-tracker/process/service"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"testing"
)

func TestAddProcessSuccessfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := mock_domain.NewMockRepository(ctrl)
	m.EXPECT().Save(gomock.Any()).Return(nil)
	processUsecase := service.NewProcessService(m)
	err := processUsecase.AddProcess(uuid.New().String(),
		"own",
		"nezasa",
		"fulltime",
		"2023-04-20", domain.WithSalary(6000, "usd"))
	if err != nil {
		t.Errorf("Error trying to add a Process %s", err.Error())
	}
}
