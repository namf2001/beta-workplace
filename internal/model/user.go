package model

import (
	"time"

	"github.com/namf2001/beta-workplace/internal/model/common"
)

// User represents a user in the system
type User struct {
	ID            int64     `json:"id,omitempty"             db:"id"`
	Name          string    `json:"name,omitempty"           db:"name"`
	Email         string    `json:"email,omitempty"          db:"email"`
	EmailVerified time.Time `json:"email_verified,omitempty" db:"email_verified"`
	Image         string    `json:"image,omitempty"          db:"image"`
	Password      string    `json:"-"                        db:"password"`
	CreatedAt     time.Time `json:"created_at,omitempty"     db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"     db:"updated_at"`
}

// Prepare auto-sets timestamps before insert/update
func (u *User) Prepare() {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	u.UpdatedAt = time.Now()
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
