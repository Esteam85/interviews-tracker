package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFirstContact(t *testing.T) {
	_, err := NewFirstContact(
		"2023-04-15",
		"Mail",
	)
	assert.NoError(t, err)
}

func TestNewFirstContactWithValidAnsweredDate(t *testing.T) {
	_, err := NewFirstContact(
		"2023-04-15",
		"Mail",
		WithAnsweredDate("2023-04-16"),
	)
	assert.NoError(t, err)
}

func TestNewFirstContactWithInValidFormatAnsweredDate(t *testing.T) {
	_, err := NewFirstContact(
		"2023-04-15",
		"Mail",
		WithAnsweredDate("2023-not-valid"),
	)
	assert.Error(t, err)
}

func TestNewFirstContactWithValidAnsweredDateFormatButBeforeThanFirstContactDate(t *testing.T) {
	_, err := NewFirstContact(
		"2023-04-15",
		"Mail",
		WithAnsweredDate("2023-04-14"),
	)
	assert.Error(t, err)
}

func TestNewFirstContactWithValidAnsweredDateFormatButBeforeEqualThanFirstContactDate(t *testing.T) {
	date := "2023-04-15"
	_, err := NewFirstContact(
		date,
		"Mail",
		WithAnsweredDate(date),
	)
	assert.NoError(t, err)
}
