package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/devdavidalonso/ecs-manager-bff/internal/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// StudentHandler coordinates HTTP endpoints for Student operations.
type StudentHandler struct {
	usecase domain.StudentUsecase
}

// NewStudentHandler creates a StudentHandler and registers endpoints on the provided multiplexer.
func NewStudentHandler(mux *http.ServeMux, us domain.StudentUsecase) {
	handler := &StudentHandler{
		usecase: us,
	}

	mux.HandleFunc("POST /api/v1/students", handler.Enroll)
	mux.HandleFunc("GET /api/v1/students/{id}", handler.GetByID)
}

// Enroll handles POST requests to enroll a new student.
func (h *StudentHandler) Enroll(w http.ResponseWriter, r *http.Request) {
	var student domain.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	err := h.usecase.EnrollStudent(r.Context(), &student)
	if err != nil {
		if errors.Is(err, domain.ErrEmptyStudentName) || errors.Is(err, domain.ErrEmptyGuardianName) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(student); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetByID handles GET requests to retrieve a student by ID.
func (h *StudentHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "missing student ID", http.StatusBadRequest)
		return
	}

	student, err := h.usecase.GetStudent(r.Context(), id)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			http.Error(w, "student not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(student); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
