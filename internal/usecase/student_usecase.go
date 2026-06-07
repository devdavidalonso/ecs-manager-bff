package usecase

import (
	"context"
	"time"

	"github.com/devdavidalonso/ecs-manager-bff/internal/domain"
)

type studentUsecase struct {
	repo domain.StudentRepository
}

// NewStudentUsecase creates a new StudentUsecase instance with dependency injection.
func NewStudentUsecase(repo domain.StudentRepository) domain.StudentUsecase {
	return &studentUsecase{
		repo: repo,
	}
}

// EnrollStudent validates student fields, determines the class group based on age,
// sets the status to pre_enrolled, and persists the record.
func (u *studentUsecase) EnrollStudent(ctx context.Context, student *domain.Student) error {
	if student.FullName == "" {
		return domain.ErrEmptyStudentName
	}
	if student.Guardian.Name == "" {
		return domain.ErrEmptyGuardianName
	}

	// Calculate age and assign class group
	age := calculateAge(student.BirthDate)
	switch {
	case age <= 3:
		student.Group = "Maternal"
	case age >= 4 && age <= 6:
		student.Group = "Jardim"
	case age >= 7 && age <= 12:
		student.Group = "Ciclos"
	default:
		student.Group = "Mocidade"
	}

	// Override status to pre_enrolled
	student.Status = domain.StatusPreEnrolled

	return u.repo.Create(ctx, student)
}

// GetStudent retrieves a student by their unique ID.
func (u *studentUsecase) GetStudent(ctx context.Context, id string) (*domain.Student, error) {
	return u.repo.GetByID(ctx, id)
}

// calculateAge calculates the calendar age based on birth date and current time.
func calculateAge(birthDate time.Time) int {
	now := time.Now()
	age := now.Year() - birthDate.Year()
	if now.Month() < birthDate.Month() || (now.Month() == birthDate.Month() && now.Day() < birthDate.Day()) {
		age--
	}
	return age
}
