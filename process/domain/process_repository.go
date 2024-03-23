package domain

type Repository interface {
	Save(process *Process) error
}
