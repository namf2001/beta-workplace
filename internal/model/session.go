package model

import "time"

// Session represents a user session
type Session struct {
	ID           int64     `json:"id,omitempty"            db:"id"`
	UserID       int64     `json:"user_id,omitempty"       db:"user_id"`
	Expires      time.Time `json:"expires,omitempty"       db:"expires"`
	SessionToken string    `json:"session_token,omitempty" db:"session_token"`
}
