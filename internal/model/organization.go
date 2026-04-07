package model

import "time"

// Organization represents a top-level organization entity
type Organization struct {
	ID          int64     `json:"id,omitempty"          db:"id"`
	Name        string    `json:"name,omitempty"        db:"name"`
	Slug        string    `json:"slug,omitempty"        db:"slug"`
	Industry    string    `json:"industry,omitempty"    db:"industry"`
	Size        int       `json:"size,omitempty"        db:"size"`
	Region      string    `json:"region,omitempty"      db:"region"`
	LogoURL     string    `json:"logo_url,omitempty"    db:"logo_url"`
	Description string    `json:"description,omitempty" db:"description"`
	CreatedBy   int64     `json:"created_by,omitempty"  db:"created_by"`
	CreatedAt   time.Time `json:"created_at,omitempty"  db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"  db:"updated_at"`
}

// Prepare auto-sets timestamps before insert/update
func (o *Organization) Prepare() {
	if o.CreatedAt.IsZero() {
		o.CreatedAt = time.Now()
	}
	o.UpdatedAt = time.Now()
}

// OrganizationMember represents a user's membership in an organization
type OrganizationMember struct {
	ID             int64     `json:"id,omitempty"              db:"id"`
	OrganizationID int64     `json:"organization_id,omitempty" db:"organization_id"`
	UserID         int64     `json:"user_id,omitempty"         db:"user_id"`
	UserName       string    `json:"user_name,omitempty"       db:"user_name"`
	UserEmail      string    `json:"user_email,omitempty"      db:"user_email"`
	UserAvatar     string    `json:"user_avatar,omitempty"     db:"user_avatar"`
	Role           OrgRole   `json:"role,omitempty"            db:"role"`
	InvitedBy      *int64    `json:"invited_by,omitempty"      db:"invited_by"`
	JoinedAt       time.Time `json:"joined_at,omitempty"       db:"joined_at"`
}

// OrganizationInvitation represents an invitation to join an organization
type OrganizationInvitation struct {
	ID             int64     `json:"id,omitempty"              db:"id"`
	OrganizationID int64     `json:"organization_id,omitempty" db:"organization_id"`
	InviteCode     string    `json:"invite_code,omitempty"     db:"invite_code"`
	Email          string    `json:"email,omitempty"           db:"email"`
	Role           OrgRole   `json:"role,omitempty"            db:"role"`
	InvitedBy      int64     `json:"invited_by,omitempty"      db:"invited_by"`
	Status         string    `json:"status,omitempty"          db:"status"`
	ExpiresAt      time.Time `json:"expires_at,omitempty"      db:"expires_at"`
	CreatedAt      time.Time `json:"created_at,omitempty"      db:"created_at"`
}
