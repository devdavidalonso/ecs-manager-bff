package domain

import (
	"context"
	"errors"
)

var (
	// ErrEmptyStudentName is returned when a student's full name is missing.
	ErrEmptyStudentName = errors.New("student full name cannot be empty")
	// ErrEmptyGuardianName is returned when the guardian's name is missing.
	ErrEmptyGuardianName = errors.New("guardian name cannot be empty")
)

// StudentUsecase defines the use cases for managing students.
type StudentUsecase interface {
	EnrollStudent(ctx context.Context, student *Student) error
	GetStudent(ctx context.Context, id string) (*Student, error)
}
