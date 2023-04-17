package domain

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Salary struct {
	Currency   Currency     `json:"currency,omitempty"`
	Amount     int          `json:"amount,omitempty"`
	SalaryType SalaryType   `json:"salaryType,omitempty"`
	Period     SalaryPeriod `json:"period,omitempty"`
}

type Currency int

const (
	usd Currency = iota
	clp
)

var currencyMap = map[string]Currency{
	"usd": usd,
	"clp": clp,
}

func ParseCurrency(s string) (Currency, error) {
	if c, ok := currencyMap[strings.ToLower(s)]; ok {
		return c, nil
	}
	return 0, fmt.Errorf("invalid currency value: %q", s)
}

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

type SalaryPeriod int

const (
	monthly SalaryPeriod = iota
	yearly
)

var salaryPeriodMap = map[string]SalaryPeriod{
	"monthly": monthly,
	"yearly":  yearly,
}

func ParseSalaryPeriod(s string) (SalaryPeriod, error) {
	if p, ok := salaryPeriodMap[strings.ToLower(s)]; ok {
		return p, nil
	}
	return 0, fmt.Errorf("invalid salary period value: %q", s)
}

type ProcessID struct {
	value string
}

type PostulationType int

const (
	own PostulationType = iota
	recruiter
	reference
)

var postulationTypeMap = map[string]PostulationType{
	"own":       own,
	"recruiter": recruiter,
	"reference": reference,
}

type Platform int

const (
	linkedin Platform = iota
	companyPortal
	getOnBoard
	compuTrabajo
)

var platformMap = map[string]Platform{
	"linkedin":      linkedin,
	"companyportal": companyPortal,
	"computrabajo":  compuTrabajo,
	"getOnBoard":    getOnBoard,
}

func ParsePlatform(s string) (Platform, error) {
	if p, ok := platformMap[strings.ToLower(s)]; ok {
		return p, nil
	}
	return 0, fmt.Errorf("invalid platform type value: %q", s)
}

func ParsePostulationType(s string) (PostulationType, error) {
	if p, ok := postulationTypeMap[strings.ToLower(s)]; ok {
		return p, nil
	}
	return 0, fmt.Errorf("invalid postulation type value: %q", s)
}

type JobType int

const (
	contract JobType = iota
	fulltime
)

var jobTypeMap = map[string]JobType{
	"contract": contract,
	"fulltime": fulltime,
}

func ParseJobType(s string) (JobType, error) {
	if p, ok := jobTypeMap[strings.ToLower(s)]; ok {
		return p, nil
	}
	return 0, fmt.Errorf("invalid job type value: %q", s)
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

func (p *Process) ProcessID() *ProcessID {
	return p.ID
}

type Repository interface {
	Save(process *Process) error
}

type Process struct {
	ID              *ProcessID      `json:"id,omitempty"`
	Platform        Platform        `json:"platform,omitempty"`
	Company         string          `json:"company,omitempty"`
	Client          string          `json:"client,omitempty"`
	Position        string          `json:"position,omitempty"`
	JobType         JobType         `json:"jobType,omitempty"`
	PostulationType PostulationType `json:"postulationType,omitempty"`
	PostulationDate time.Time       `json:"postulationDate"`
	FirstContact    *FirstContact   `json:"firstContact,omitempty"`
	Salary          *Salary         `json:"salary,omitempty"`
}

func NewProcess(id,
	postulationType,
	platform,
	company,
	position,
	jobType string,
	options ...func(p *Process) error) (*Process, error) {

	processID, err := NewProcessID(id)
	if err != nil {
		return &Process{}, err
	}

	pType, err := ParsePostulationType(postulationType)
	if err != nil {
		return &Process{}, err
	}

	jType, err := ParseJobType(jobType)
	if err != nil {
		return &Process{}, err
	}

	p, err := ParsePlatform(platform)
	if err != nil {
		return &Process{}, err
	}
	process := &Process{
		ID:              processID,
		PostulationType: pType,
		Position:        position,
		Company:         company,
		JobType:         jType,
		PostulationDate: time.Now(),
		Platform:        p,
	}

	for _, o := range options {
		err = o(process)
		if err != nil {
			return &Process{}, err
		}
	}
	return process, nil
}

func WithSalary(amount int, currency, salaryType, period string) func(*Process) error {
	return func(p *Process) error {
		c, err := ParseCurrency(currency)
		if err != nil {
			return fmt.Errorf("invalid currency value: %q", currency)
		}
		s, err := ParseSalaryType(salaryType)
		if err != nil {
			return err
		}
		sP, err := ParseSalaryPeriod(period)
		if err != nil {
			return err
		}
		salary := &Salary{
			Amount:     amount,
			Currency:   c,
			SalaryType: s,
			Period:     sP,
		}
		p.Salary = salary
		return nil
	}
}

func WithClient(client string) func(*Process) error {
	return func(p *Process) error {
		p.Client = client
		return nil
	}
}

func WithFirstContact(f *FirstContact) func(*Process) error {
	return func(p *Process) error {
		p.FirstContact = f
		return nil
	}
}
