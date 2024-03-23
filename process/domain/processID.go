package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type ProcessID struct {
	value string
}

func NewProcessID(value string) (*ProcessID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return &ProcessID{}, fmt.Errorf("invalid uuid value: %s", err.Error())
	}
	return &ProcessID{
		value: v.String(),
	}, nil
}

func (pId *ProcessID) String() string {
	return pId.value
}
