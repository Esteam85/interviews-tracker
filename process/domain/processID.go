package domain

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var ErrInvalidProcessID = errors.New("invalid process id probably not uuid format")

type ProcessID struct {
	value string
}

func NewProcessID(value string) (*ProcessID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return &ProcessID{}, fmt.Errorf("invalid process id value: %s, %w", value, ErrInvalidProcessID)
	}
	return &ProcessID{
		value: v.String(),
	}, nil
}

func (pId *ProcessID) String() string {
	return pId.value
}
