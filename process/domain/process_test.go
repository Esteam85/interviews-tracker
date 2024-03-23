package domain

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
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
		WithSalary(&SalaryAsPrimitives{Amount: 6000, Currency: "usd", SalaryType: "gross", SalaryPeriod: "monthly"}))
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
		WithSalary(&SalaryAsPrimitives{Amount: 6000, Currency: "usd", SalaryType: "gross", SalaryPeriod: "monthly"}),
		WithClient("client"))
	assert.NoError(t, err)
}

func TestAddProcessWithFirstContactDate(t *testing.T) {
	id, _ := uuid.NewUUID()
	_, err := NewProcess(
		id.String(),
		"Own",
		"Linkedin",
		"Esteam",
		"Dev",
		"contract",
		WithFirstContact(&FirstContactAsPrimitives{"2023-04-15", "Mail", "2023-04-24"}))
	assert.NoError(t, err)
}
