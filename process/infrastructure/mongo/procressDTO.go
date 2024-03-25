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
	pAsPrimitives := p.Process
	return domain.NewProcess(
		pAsPrimitives.ProcessID,
		pAsPrimitives.PostulationType,
		pAsPrimitives.Platform,
		pAsPrimitives.Company,
		pAsPrimitives.Position,
		pAsPrimitives.JobType,
		domain.WithFirstContact(pAsPrimitives.FirstContact),
		domain.WithSalary(pAsPrimitives.Salary),
		domain.WithClient(pAsPrimitives.Client),
	)
}