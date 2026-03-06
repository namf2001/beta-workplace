package model

import (
	"time"

	"github.com/namf2001/beta-workplace/internal/model/common"
)

// User represents a user in the system
type User struct {
	ID            int64      `json:"id" db:"id"`
	Email         string     `json:"email" db:"email"`
	Name          string     `json:"name" db:"name"`
	EmailVerified *time.Time `json:"emailVerified" db:"emailVerified"`
	Image         string     `json:"image" db:"image"`
	Password      string     `json:"-" db:"password"` // Stored in users now
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at" db:"updated_at"`
}

// Validate validates user data
func (u *User) Validate() error {
	if u.Email == "" {
		return common.ErrInvalidEmail
	}
	if u.Name == "" {
		return common.ErrInvalidName
	}
	return nil
}
