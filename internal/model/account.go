package model

import "time"

// Account represents an OAuth account connection
type Account struct {
	ID                int64     `json:"id,omitempty"                  db:"id"`
	UserID            int64     `json:"user_id,omitempty"             db:"user_id"`
	Type              string    `json:"type,omitempty"                db:"type"`
	Provider          Provider  `json:"provider,omitempty"            db:"provider"`
	ProviderAccountID string    `json:"provider_account_id,omitempty" db:"provider_account_id"`
	RefreshToken      string    `json:"refresh_token,omitempty"       db:"refresh_token"`
	AccessToken       string    `json:"access_token,omitempty"        db:"access_token"`
	ExpiresAt         int64     `json:"expires_at,omitempty"          db:"expires_at"`
	IDToken           string    `json:"id_token,omitempty"            db:"id_token"`
	Scope             string    `json:"scope,omitempty"               db:"scope"`
	SessionState      string    `json:"session_state,omitempty"       db:"session_state"`
	TokenType         string    `json:"token_type,omitempty"          db:"token_type"`
	CreatedAt         time.Time `json:"created_at,omitempty"          db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at,omitempty"          db:"updated_at"`
}

// Prepare auto-sets timestamps before insert/update
func (a *Account) Prepare() {
	if a.CreatedAt.IsZero() {
		a.CreatedAt = time.Now()
	}
	a.UpdatedAt = time.Now()
}
