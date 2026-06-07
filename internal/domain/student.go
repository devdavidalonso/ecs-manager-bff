package domain

import "time"

// StudentStatus represents the enrollment status of a student.
type StudentStatus string

const (
	StatusPreEnrolled StudentStatus = "pre_enrolled"
	StatusConfirmed   StudentStatus = "confirmed"
	StatusCancelled   StudentStatus = "cancelled"
	StatusWaitingList StudentStatus = "waiting_list"
)

// Guardian represents the contact details of the student's parent or legal guardian.
type Guardian struct {
	Name  string `json:"name" firestore:"name"`
	Phone string `json:"phone" firestore:"phone"`
	Email string `json:"email" firestore:"email"`
}

// Student represents a child or teenager registered in the evangelization program.
type Student struct {
	ID                 string        `json:"id" firestore:"-"`
	FullName           string        `json:"full_name" firestore:"full_name"`
	BirthDate          time.Time     `json:"birth_date" firestore:"birth_date"`
	HowKnewUs          string        `json:"how_knew_us" firestore:"how_knew_us"`
	FamilyMembersCount int           `json:"family_members_count" firestore:"family_members_count"`
	Guardian           Guardian      `json:"guardian" firestore:"guardian"`
	Status             StudentStatus `json:"status" firestore:"status"`
	Group              string        `json:"group" firestore:"group"`
	CreatedAt          time.Time     `json:"created_at" firestore:"created_at"`
	UpdatedAt          time.Time     `json:"updated_at" firestore:"updated_at"`
}
