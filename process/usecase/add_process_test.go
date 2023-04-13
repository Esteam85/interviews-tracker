package usecase_test

import (
	"github.com/esteam85/interviews-tracker/process/domain"
	"github.com/esteam85/interviews-tracker/process/usecase"
	"testing"
)

func TestAddProcessSuccessfully(t *testing.T) {

	// arrange
	processUsecase := usecase.NewProcessUsecase()
	// act
	err := processUsecase.AddProcess("recruiter",
		"nezasa",
		"nezasa",
		"full-time",
		"09-03-2023", domain.WithSalary(6000, "usd"))
	if err != nil {
		t.Errorf("Error trying to add a Process %s", err.Error())
	}
	// assert
}
