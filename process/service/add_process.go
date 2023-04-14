package service

import "github.com/esteam85/interviews-tracker/process/domain"

type ProcessService struct {
	repository domain.Repository
}

func NewProcessService(r domain.Repository) *ProcessService {
	return &ProcessService{
		repository: r,
	}
}
func (p *ProcessService) AddProcess(id, postulationType, company, jobType, firstContactDate string, options ...func(*domain.Process) error) error {

	process, err := domain.NewProcess(id, postulationType, company, jobType, firstContactDate, options...)
	if err != nil {
		return err
	}

	err = p.repository.Save(process)
	if err != nil {
		return err
	}

	return nil
}
