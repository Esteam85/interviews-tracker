package usecase

import "github.com/esteam85/interviews-tracker/process/domain"

type ProcessUsecase struct {
}

func NewProcessUsecase() *ProcessUsecase {
	return &ProcessUsecase{}
}
func (p *ProcessUsecase) AddProcess(id, postulationType, company, jobType, firstContactDate string, options ...func(*domain.Process) error) error {

	_, _ = domain.NewProcess(id, postulationType, company, jobType, firstContactDate, options...)

	return nil
}
