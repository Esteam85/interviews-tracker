package domain

import (
	"github.com/google/uuid"
	"testing"
)

func TestAddProcessWithoutOptions(t *testing.T) {

	id, _ := uuid.NewUUID()
	_, err := NewProcess(
		id.String(),
		"Own",
		"Esteam",
		"contract",
		"2023-04-12")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestAddProcessWithSalary(t *testing.T) {

	id, _ := uuid.NewUUID()
	_, err := NewProcess(
		id.String(),
		"Own",
		"Esteam",
		"contract",
		"2023-04-12",
		WithSalary(6000, "usd"))
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestAddProcessWithSalaryAndClient(t *testing.T) {
	id, _ := uuid.NewUUID()
	_, err := NewProcess(
		id.String(),
		"Own",
		"Esteam",
		"contract",
		"2023-04-12",
		WithSalary(6000, "usd"),
		WithClient("client"))
	if err != nil {
		t.Errorf(err.Error())
	}
}
