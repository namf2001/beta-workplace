package organizations

import "time"

// ─── Organization Responses ────────────────────────────────────────────────

// OrganizationDetailResponse is the handler-level DTO for a single organization
type OrganizationDetailResponse struct {
	ID            int64     `json:"id"`
	Name          string    `json:"name"`
	Industry      string    `json:"industry"`
	Size          int       `json:"size"`
	Region        string    `json:"region"`
	Avatar        string    `json:"avatar"`
	Description   string    `json:"description"`
	AdminCount    int       `json:"admin_count"`
	MemberCount   int       `json:"member_count"`
	SubAdminCount int       `json:"sub_admin_count"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// UpdateOrganizationResponse is the handler-level DTO after updating an organization
type UpdateOrganizationResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// OrganizationSummaryResponse is the brief org info nested in the list
type OrganizationSummaryResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Industry string `json:"industry"`
}

// OrganizationWithRoleResponse is one item in the user organizations list
type OrganizationWithRoleResponse struct {
	Organization OrganizationSummaryResponse `json:"organization"`
	Role         string                      `json:"role"`
	JoinedAt     string                      `json:"joined_at"`
}

// UserOrganizationsResponse is the paginated list of user organizations
type UserOrganizationsResponse struct {
	Organizations []OrganizationWithRoleResponse `json:"organizations"`
	Total         int64                          `json:"total"`
}

// ─── Member Responses ─────────────────────────────────────────────────────

// MemberDetailResponse is the handler-level DTO for a single member
type MemberDetailResponse struct {
	ID             int64     `json:"id"`
	OrganizationID int64     `json:"organization_id"`
	UserID         int64     `json:"user_id"`
	UserName       string    `json:"user_name"`
	UserEmail      string    `json:"user_email"`
	UserAvatar     string    `json:"user_avatar"`
	Role           string    `json:"role"`
	InvitedBy      *int64    `json:"invited_by,omitempty"`
	JoinedAt       time.Time `json:"joined_at"`
}

// JoinOrganizationResponse is the handler-level DTO returned after joining an organization
type JoinOrganizationResponse struct {
	ID             int64     `json:"id"`
	OrganizationID int64     `json:"organization_id"`
	UserID         int64     `json:"user_id"`
	UserName       string    `json:"user_name"`
	UserEmail      string    `json:"user_email"`
	UserAvatar     string    `json:"user_avatar"`
	Role           string    `json:"role"`
	InvitedBy      *int64    `json:"invited_by,omitempty"`
	JoinedAt       time.Time `json:"joined_at"`
}

// ─── Invitation Responses ─────────────────────────────────────────────────

// InvitationDetailResponse is the handler-level DTO for a created invitation
type InvitationDetailResponse struct {
	ID             int64     `json:"id"`
	OrganizationID int64     `json:"organization_id"`
	InviteCode     string    `json:"invite_code"`
	Email          string    `json:"email"`
	ExpiresAt      time.Time `json:"expires_at"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
}
