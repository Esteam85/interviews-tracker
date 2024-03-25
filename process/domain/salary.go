package domain

import "fmt"

type Salary struct {
	Currency   Currency     `json:"currency,omitempty"`
	Amount     int          `json:"amount,omitempty"`
	SalaryType SalaryType   `json:"salaryType,omitempty"`
	Period     SalaryPeriod `json:"period,omitempty"`
}

func NewSalary(sAsPrimitives *SalaryAsPrimitives) (*Salary, error) {
	c, err := ParseCurrency(sAsPrimitives.Currency)
	if err != nil {
		return nil, fmt.Errorf("invalid currency value: %s", sAsPrimitives.Currency)
	}
	s, err := ParseSalaryType(sAsPrimitives.SalaryType)
	if err != nil {
		return nil, err
	}
	sP, err := ParseSalaryPeriod(sAsPrimitives.SalaryPeriod)
	if err != nil {
		return nil, err
	}
	return &Salary{
		Amount:     sAsPrimitives.Amount,
		Currency:   c,
		SalaryType: s,
		Period:     sP,
	}, nil
}
