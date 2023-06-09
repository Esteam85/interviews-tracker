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

func TestAddProcessWithFirstContactDate(t *testing.T) {
	id, _ := uuid.NewUUID()
	firstContact, err := NewFirstContact(
		"2023-04-15",
		"Mail",
		WithAnsweredDate("2023-04-24"),
	)
	assert.NoError(t, err)
	_, err = NewProcess(
		id.String(),
		"Own",
		"Linkedin",
		"Esteam",
		"Dev",
		"contract",
		WithFirstContact(
			firstContact))
	assert.NoError(t, err)
}
