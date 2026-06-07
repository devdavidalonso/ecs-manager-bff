package domain

import "context"

// StudentRepository defines the data access contract for Student entities.
type StudentRepository interface {
	Create(ctx context.Context, student *Student) error
	GetByID(ctx context.Context, id string) (*Student, error)
}
