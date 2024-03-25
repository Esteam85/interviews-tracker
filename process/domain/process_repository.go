package domain

import "context"

type ProcessRepository interface {
	Save(ctx context.Context, process *Process) error
	GetAll(ctx context.Context) ([]Process, error)
}
