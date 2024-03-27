package mongo

import (
	"github.com/esteam85/interviews-tracker/process/domain"
)

type ProcessDTO struct {
	ID      string                      `bson:"_id,omitempty"`
	Process *domain.ProcessAsPrimitives `bson:"process"`
}

func fromPrimitives(p *domain.ProcessAsPrimitives) *ProcessDTO {
	return &ProcessDTO{
		ID:      p.ProcessID,
		Process: p,
	}
}

func (p ProcessDTO) ToProcess() (*domain.Process, error) {
	return domain.NewProcess(
		p.Process.ProcessID,
		p.Process.PostulationType,
		p.Process.Platform,
		p.Process.Company,
		p.Process.Position,
		p.Process.JobType,
		domain.WithFirstContact(p.Process.FirstContact),
		domain.WithSalary(p.Process.Salary),
		domain.WithClient(p.Process.Client),
	)
}
