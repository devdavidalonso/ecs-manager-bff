package domain

import "time"

// VolunteerStatus represents the active or inactive state of a volunteer.
type VolunteerStatus string

const (
	StatusActive   VolunteerStatus = "active"
	StatusInactive VolunteerStatus = "inactive"
)

// Volunteer represents a collaborator in the evangelization institution.
type Volunteer struct {
	ID              string          `json:"id" firestore:"-"`
	FullName        string          `json:"full_name" firestore:"full_name"`
	Phone           string          `json:"phone" firestore:"phone"`
	Email           string          `json:"email" firestore:"email"`
	AreasOfInterest []string        `json:"areas_of_interest" firestore:"areas_of_interest"`
	Status          VolunteerStatus `json:"status" firestore:"status"`
	CreatedAt       time.Time       `json:"created_at" firestore:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at" firestore:"updated_at"`
}
