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
		"Linkedin",
		"Esteam",
		"Dev",
		"contract")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestAddProcessWithSalary(t *testing.T) {

	id, _ := uuid.NewUUID()
	_, err := NewProcess(
		id.String(),
		"Own",
		"Linkedin",
		"Esteam",
		"Dev",
		"contract",
		WithSalary(6000, "usd", "gross", "monthly"))
	assert.NoError(t, err)
}

func TestAddProcessWithSalaryAndClient(t *testing.T) {
	id, _ := uuid.NewUUID()
	_, err := NewProcess(
		id.String(),
		"Own",
		"Linkedin",
		"Esteam",
		"Dev",
		"contract",
		WithSalary(6000, "usd", "gross", "monthly"),
		WithClient("client"))
	assert.NoError(t, err)
}
