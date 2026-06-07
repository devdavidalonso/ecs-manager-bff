package firestore

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/devdavidalonso/ecs-manager-bff/internal/domain"
)

type studentRepository struct {
	client *firestore.Client
}

// NewStudentRepository creates a new instance of domain.StudentRepository using Cloud Firestore.
func NewStudentRepository(client *firestore.Client) domain.StudentRepository {
	return &studentRepository{
		client: client,
	}
}

// Create persists a new Student in the Firestore "students" collection.
// It sets the created_at and updated_at timestamps before saving.
func (r *studentRepository) Create(ctx context.Context, student *domain.Student) error {
	now := time.Now()
	student.CreatedAt = now
	student.UpdatedAt = now

	var docRef *firestore.DocumentRef
	if student.ID == "" {
		docRef = r.client.Collection("students").NewDoc()
		student.ID = docRef.ID
	} else {
		docRef = r.client.Collection("students").Doc(student.ID)
	}

	_, err := docRef.Set(ctx, student)
	return err
}

// GetByID retrieves a Student document by its ID and decodes it.
// It maps the document ID to the Student.ID field.
func (r *studentRepository) GetByID(ctx context.Context, id string) (*domain.Student, error) {
	docSnap, err := r.client.Collection("students").Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}

	var student domain.Student
	if err := docSnap.DataTo(&student); err != nil {
		return nil, err
	}

	student.ID = docSnap.Ref.ID
	return &student, nil
}
