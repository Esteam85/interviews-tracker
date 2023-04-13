package domain

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Salary struct {
	currency Currency
	amount   int
}

type Currency int

const (
	USD Currency = iota
	CLP
)

var currencyMap = map[string]Currency{
	"usd": USD,
	"clp": CLP,
}

func ParseCurrency(s string) (Currency, error) {
	if c, ok := currencyMap[strings.ToLower(s)]; ok {
		return c, nil
	}
	return 0, fmt.Errorf("invalid currency value: %q", s)
}

type ProcessID struct {
	value string
}

type PostulationType int

const (
	Own PostulationType = iota
	Recruiter
	Reference
)

var postulationTypeMap = map[string]PostulationType{
	"own":       Own,
	"recruiter": Recruiter,
	"reference": Reference,
}

func ParsePostulationType(s string) (PostulationType, error) {
	if p, ok := postulationTypeMap[strings.ToLower(s)]; ok {
		return p, nil
	}
	return 0, fmt.Errorf("invalid postulation type value: %q", s)
}

type JobType int

const (
	Contract JobType = iota
	Fulltime
)

var jobTypeMap = map[string]JobType{
	"contract": Contract,
	"fulltime": Fulltime,
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

func (p *Process) ID() *ProcessID {
	return p.id
}

type Process struct {
	id               *ProcessID
	postulationType  PostulationType
	company          string
	client           string
	jobType          JobType
	postulationDate  time.Time
	firstContactDate time.Time
	salary           *Salary
}

func NewProcess(id, pType, company,
	jType,
	fCDate string, options ...func(p *Process) error) (*Process, error) {

	processID, err := NewProcessID(id)
	if err != nil {
		return &Process{}, err
	}

	postulationType, err := ParsePostulationType(pType)
	if err != nil {
		return &Process{}, err
	}

	jobType, err := ParseJobType(jType)
	if err != nil {
		return &Process{}, err
	}

	firstContactDate, err := time.Parse("2006-01-02", fCDate)
	if err != nil {
		return &Process{}, err
	}

	process := &Process{
		id:               processID,
		postulationType:  postulationType,
		company:          company,
		jobType:          jobType,
		postulationDate:  time.Now(),
		firstContactDate: firstContactDate,
	}

	for _, o := range options {
		err = o(process)
		if err != nil {
			return &Process{}, err
		}
	}
	return process, nil
}

func WithSalary(amount int, currency string) func(*Process) error {
	return func(p *Process) error {
		c, err := ParseCurrency(currency)
		if err != nil {
			return fmt.Errorf("invalid currency value: %q", currency)
		}
		salary := &Salary{
			amount:   amount,
			currency: c,
		}
		p.salary = salary
		return nil
	}
}

func WithClient(client string) func(*Process) error {
	return func(p *Process) error {
		p.client = client
		return nil
	}
}