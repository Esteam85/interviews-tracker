package domain

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
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
	assert.NoError(t, err)
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
	assert.NoError(t, err)
}
