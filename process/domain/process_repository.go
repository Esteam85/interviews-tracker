package domain

import "context"

type ProcessRepository interface {
	Save(ctx context.Context, process *Process) error
}
