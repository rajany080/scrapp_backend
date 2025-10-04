package schemas

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

// Valid user roles
const (
	RoleGeneral = "General"
	RoleAdmin   = "Admin"
	RoleDealer  = "Dealer"
)

var ValidRoles = []string{RoleGeneral, RoleAdmin, RoleDealer}

type CreateUserSchema struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone" binding:"required"`
	Password  string `json:"password" binding:"required,min=8"`
	About     string `json:"about"` // optional
	Role      string `json:"role" binding:"required"`
}

type GetUserResponse struct {
	ID        uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	FirstName string    `json:"first_name" example:"John"`
	LastName  string    `json:"last_name" example:"Doe"`
	Email     string    `json:"email" example:"john.doe@example.com"`
	Phone     string    `json:"phone" example:"+1234567890"`
	About     string    `json:"about" example:"Software developer"`
	Role      string    `json:"role" example:"General"`
	CreatedAt time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}

// Validate performs custom validation for CreateUserSchema
func (s *CreateUserSchema) Validate() error {
	// Check if role is valid
	roleValid := false
	for _, validRole := range ValidRoles {
		if s.Role == validRole {
			roleValid = true
			break
		}
	}
	if !roleValid {
		return &ValidationError{
			Field:   "role",
			Message: "must be one of: " + strings.Join(ValidRoles, ", "),
		}
	}

	// Validate password strength
	if len(s.Password) < 8 {
		return &ValidationError{
			Field:   "password",
			Message: "must be at least 8 characters long",
		}
	}

	return nil
}

type LoginSchema struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Message string          `json:"message" example:"Login successful"`
	User    GetUserResponse `json:"user"`
}

// ValidationError represents a validation error
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return e.Field + " " + e.Message
}
