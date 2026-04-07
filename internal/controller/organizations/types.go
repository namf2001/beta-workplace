package organizations

import (
	"time"

	"github.com/namf2001/beta-workplace/internal/model"
)

// ─── Input Types ──────────────────────────────────────────────────────────────

type CreateOrganizationInput struct {
	Name        string `json:"name"`
	Industry    string `json:"industry"`
	Region      string `json:"region"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
	Size        int    `json:"size"`
}

type UpdateOrganizationInput struct {
	Name        string `json:"name"`
	Industry    string `json:"industry"`
	Size        int    `json:"size"`
	Region      string `json:"region"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
}

type InviteMemberInput struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}

type UpdateMemberRoleInput struct {
	UserID int64  `json:"user_id"`
	Role   string `json:"role"`
}

// ─── Output Types ────────────────────────────────────────────────────────────

// OrganizationResponse is returned for create/get-by-id operations
type OrganizationResponse struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	Industry     string    `json:"industry"`
	Size         int       `json:"size"`
	Region       string    `json:"region"`
	Avatar       string    `json:"avatar"`
	Description  string    `json:"description"`
	AdminCount   int       `json:"admin_count"`
	MemberCount  int       `json:"member_count"`
	SubAdminCount int      `json:"sub_admin_count"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// UpdateOrganizationResponse is returned after updating an organization
type UpdateOrganizationResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// OrganizationSummary is the brief org info nested inside list items
type OrganizationSummary struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Industry string `json:"industry"`
}

// OrganizationWithRoleResponse is a single item in the user organizations list
type OrganizationWithRoleResponse struct {
	Organization OrganizationSummary `json:"organization"`
	Role         model.OrgRole       `json:"role"`
	JoinedAt     string              `json:"joined_at"`
}

// OrganizationListResponse is returned by GetUserOrganizations
type OrganizationListResponse struct {
	Organizations []OrganizationWithRoleResponse `json:"organizations"`
	Total         int64                          `json:"total"`
}

// MemberResponse is a single member entry in GetMembers response
type MemberResponse struct {
	ID             int64       `json:"id"`
	OrganizationID int64       `json:"organization_id"`
	UserID         int64       `json:"user_id"`
	UserName       string      `json:"user_name"`
	UserEmail      string      `json:"user_email"`
	UserAvatar     string      `json:"user_avatar"`
	Role           model.OrgRole `json:"role"`
	InvitedBy      *int64      `json:"invited_by,omitempty"`
	JoinedAt       time.Time   `json:"joined_at"`
}

// InvitationResponse is returned after creating an invitation
type InvitationResponse struct {
	ID             int64     `json:"id"`
	OrganizationID int64     `json:"organization_id"`
	InviteCode     string    `json:"invite_code"`
	Email          string    `json:"email"`
	ExpiresAt      time.Time `json:"expires_at"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
}

// JoinOrganizationResponse is returned after joining an organization
type JoinOrganizationResponse struct {
	ID             int64       `json:"id"`
	OrganizationID int64       `json:"organization_id"`
	UserID         int64       `json:"user_id"`
	UserName       string      `json:"user_name"`
	UserEmail      string      `json:"user_email"`
	UserAvatar     string      `json:"user_avatar"`
	Role           model.OrgRole `json:"role"`
	InvitedBy      *int64      `json:"invited_by,omitempty"`
	JoinedAt       time.Time   `json:"joined_at"`
}
