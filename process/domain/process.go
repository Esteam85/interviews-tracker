package domain

import (
	"time"
)

type ProcessAsPrimitives struct {
	ProcessID       string                    `json:"process_id"`
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
	processID       *ProcessID
	postulationType PostulationType
	platform        Platform
	company         string
	position        string
	jobType         JobType
	client          string
	firstContact    *FirstContact
	salary          *Salary
	postulationDate time.Time
}

func NewProcess(
	id,
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
		processID:       processID,
		postulationType: pType,
		position:        position,
		company:         company,
		jobType:         jType,
		postulationDate: time.Now(),
		platform:        p,
	}

	for _, o := range options {
		if o == nil {
			continue
		}
		err = o(process)
		if err != nil {
			return &Process{}, err
		}
	}
	return process, nil
}

func (p *Process) ProcessID() *ProcessID {
	return p.processID
}

func (p *Process) ToPrimitives() *ProcessAsPrimitives {

	processAsPrimitives := &ProcessAsPrimitives{
		ProcessID:       p.ProcessID().String(),
		Platform:        p.platform.String(),
		Company:         p.company,
		Client:          p.client,
		Position:        p.position,
		JobType:         p.jobType.String(),
		PostulationType: p.postulationType.String(),
	}

	if p.firstContact != nil {
		processAsPrimitives.FirstContact = &FirstContactAsPrimitives{
			ContactDate:  p.firstContact.ContactDate.String(),
			Channel:      p.firstContact.Channel.String(),
			AnsweredDate: p.firstContact.AnsweredDate.String(),
		}
	}

	if p.salary != nil {
		processAsPrimitives.Salary = &SalaryAsPrimitives{
			Amount:       p.salary.Amount,
			Currency:     p.salary.Currency.String(),
			SalaryType:   p.salary.SalaryType.String(),
			SalaryPeriod: p.salary.Period.String(),
		}
	}
	return processAsPrimitives
}

type ProcessOptions func(*Process) error

func WithSalary(sAsPrimitives *SalaryAsPrimitives) func(*Process) error {
	if sAsPrimitives == nil {
		return nil
	}
	return func(p *Process) error {
		salary, err := NewSalary(sAsPrimitives)
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

func WithFirstContact(fc *FirstContactAsPrimitives) func(*Process) error {
	if fc == nil {
		return nil
	}
	return func(p *Process) error {
		firstContact, err := NewFirstContact(fc.ContactDate, fc.Channel, WithAnsweredDate(fc.AnsweredDate))
		if err != nil {
			return err
		}
		p.firstContact = firstContact
		return nil
	}
}
