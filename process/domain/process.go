package domain

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Salary struct {
	currency   Currency
	amount     int
	salaryType SalaryType
	period     SalaryPeriod
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

func (p *Process) ID() *ProcessID {
	return p.id
}

type Repository interface {
	Save(process *Process) error
}

type Process struct {
	id              *ProcessID
	platform        Platform
	company         string
	client          string
	position        string
	jobType         JobType
	postulationType PostulationType
	postulationDate time.Time
	firstContact    *FirstContact
	salary          *Salary
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
		id:              processID,
		postulationType: pType,
		position:        position,
		company:         company,
		jobType:         jType,
		postulationDate: time.Now(),
		platform:        p,
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
			amount:     amount,
			currency:   c,
			salaryType: s,
			period:     sP,
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

func WithFirstContact(f *FirstContact) func(*Process) error {
	return func(p *Process) error {
		p.firstContact = f
		return nil
	}
}
