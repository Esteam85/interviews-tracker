package domain

import (
	"fmt"
	"strings"
)

type SalaryPeriod int

const (
	monthly SalaryPeriod = iota
	yearly
)

var salaryPeriodMap = map[string]SalaryPeriod{
	"monthly": monthly,
	"yearly":  yearly,
}
var invertSalaryPeriodMap = map[SalaryPeriod]string{
	monthly: "monthly",
	yearly:  "yearly",
}

func ParseSalaryPeriod(s string) (SalaryPeriod, error) {
	if p, ok := salaryPeriodMap[strings.ToLower(s)]; ok {
		return p, nil
	}
	return 0, fmt.Errorf("invalid salary period value: %q", s)
}

func (s SalaryPeriod) String() string {
	return invertSalaryPeriodMap[s]
}
