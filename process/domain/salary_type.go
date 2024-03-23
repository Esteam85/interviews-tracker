package domain

import (
	"fmt"
	"strings"
)

type SalaryType int

const (
	gross SalaryType = iota
	net
)

var salaryTypeMap = map[string]SalaryType{
	"gross": gross,
	"net":   net,
}

func ParseSalaryType(s string) (SalaryType, error) {
	if c, ok := salaryTypeMap[strings.ToLower(s)]; ok {
		return c, nil
	}
	return 0, fmt.Errorf("invalid salary type value: %q", s)
}
