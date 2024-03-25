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
