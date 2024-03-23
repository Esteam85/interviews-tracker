package domain

type Salary struct {
	Currency   Currency     `json:"currency,omitempty"`
	Amount     int          `json:"amount,omitempty"`
	SalaryType SalaryType   `json:"salaryType,omitempty"`
	Period     SalaryPeriod `json:"period,omitempty"`
}
