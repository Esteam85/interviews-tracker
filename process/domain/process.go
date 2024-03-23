package domain

import (
	"time"
)

type ProcessAsPrimitives struct {
	Id              string                    `json:"id"`
	Platform        string                    `json:"platform"`
	Company         string                    `json:"company"`
	Client          string                    `json:"client"`
	Position        string                    `json:"position"`
	JobType         string                    `json:"job_type"`
	PostulationType string                    `json:"postulation_type"`
	FirstContact    *FirstContactAsPrimitives `json:"fist_contact"`
	Salary          *SalaryAsPrimitives       `json:"salary"`
}

type FirstContactAsPrimitives struct {
	ContactDate  string `json:"contact_date"`
	Channel      string `json:"channel"`
	AnsweredDate string `json:"answered_date"`
}

type SalaryAsPrimitives struct {
	Amount       int    `json:"amount"`
	Currency     string `json:"currency"`
	SalaryType   string `json:"salary_type"`
	SalaryPeriod string `json:"salary_period"`
}

type Process struct {
	id              *ProcessID
	platform        Platform
	company         string
	client          string
	position        string
	jobType         JobType
	firstContact    *FirstContact
	salary          *Salary
	postulationType PostulationType
	postulationDate time.Time
}

func NewProcess(id,
	postulationType,
	platform,
	company,
	position,
	jobType string,
	options ...ProcessOptions) (*Process, error) {

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

func (p *Process) ProcessID() *ProcessID {
	return p.id
}

type ProcessOptions func(*Process) error

func WithSalary(amount int, currency, salaryType, period string) func(*Process) error {
	return func(p *Process) error {
		salary, err := NewSalary(amount, currency, salaryType, period)
		if err != nil {
			return err
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

func WithFirstContact(date, channel string, options ...func(f *FirstContact) error) func(*Process) error {
	return func(p *Process) error {
		firstContact, err := NewFirstContact(date, channel, options...)
		if err != nil {
			return err
		}
		p.firstContact = firstContact
		return nil
	}
}
