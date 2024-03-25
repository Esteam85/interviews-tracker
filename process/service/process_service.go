package service

import (
	"context"

	"github.com/esteam85/interviews-tracker/process/domain"
)

type ProcessService struct {
	repository domain.ProcessRepository
}

func NewProcessService(r domain.ProcessRepository) *ProcessService {
	return &ProcessService{
		repository: r,
	}
}

func (p *ProcessService) GetAllProcesses(ctx context.Context) ([]domain.Process, error) {
	return p.repository.GetAll(ctx)
}

func (p *ProcessService) AddProcess(ctx context.Context, id, postulationType, platform, company, position, jobType string, options ...domain.ProcessOptions) error {
	process, err := domain.NewProcess(id, postulationType, platform, company, position, jobType, options...)
	if err != nil {
		return err
	}
	return p.repository.Save(ctx, process)
}
