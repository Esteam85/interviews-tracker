package domain

import "fmt"

type Salary struct {
	Currency   Currency     `json:"currency,omitempty"`
	Amount     int          `json:"amount,omitempty"`
	SalaryType SalaryType   `json:"salaryType,omitempty"`
	Period     SalaryPeriod `json:"period,omitempty"`
}

func NewSalary(amount int, currency, salaryType, period string) (*Salary, error) {
	c, err := ParseCurrency(currency)
	if err != nil {
		return nil, fmt.Errorf("invalid currency value: %s", currency)
	}
	s, err := ParseSalaryType(salaryType)
	if err != nil {
		return nil, err
	}
	sP, err := ParseSalaryPeriod(period)
	if err != nil {
		return nil, err
	}
	return &Salary{
		Amount:     amount,
		Currency:   c,
		SalaryType: s,
		Period:     sP,
	}, nil
}
