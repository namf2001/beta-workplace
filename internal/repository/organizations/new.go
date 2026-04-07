package organizations

import (
	"context"

	"github.com/namf2001/beta-workplace/internal/model"
	"github.com/namf2001/beta-workplace/internal/repository/db/pg"
)

type Repository interface {
	// Create creates a new organization
	Create(ctx context.Context, org model.Organization) (model.Organization, error)

	// GetByID retrieves an organization by ID
	GetByID(ctx context.Context, id int64) (model.Organization, error)

	// GetByUserID retrieves all organizations where user is a member
	GetByUserID(ctx context.Context, userID int64, page, pageSize int) ([]OrganizationWithRole, int64, error)

	// Update updates an existing organization
	Update(ctx context.Context, org model.Organization) error

	// Delete deletes an organization by ID
	Delete(ctx context.Context, id int64) error

	// AddMember adds a member to organization
	AddMember(ctx context.Context, member model.OrganizationMember) (model.OrganizationMember, error)

	// GetMember gets a specific member
	GetMember(ctx context.Context, orgID, userID int64) (model.OrganizationMember, error)

	// GetMembers gets all members of an organization
	GetMembers(ctx context.Context, orgID int64) ([]model.OrganizationMember, error)

	// UpdateMemberRole updates member role
	UpdateMemberRole(ctx context.Context, orgID, userID int64, role model.OrgRole) error

	// RemoveMember removes a member from organization
	RemoveMember(ctx context.Context, orgID, userID int64) error

	// GetMemberCounts gets counts of members by role
	GetMemberCounts(ctx context.Context, orgID int64) (admin, subAdmin, member int, err error)

	// CreateInvitation creates a new organization invitation
	CreateInvitation(ctx context.Context, inv model.OrganizationInvitation) (model.OrganizationInvitation, error)

	// GetInvitationByCode retrieves an invitation by invite code
	GetInvitationByCode(ctx context.Context, inviteCode string) (model.OrganizationInvitation, error)

	// UseInvitation marks an invitation as accepted
	UseInvitation(ctx context.Context, inviteCode string) error
}

type OrganizationWithRole struct {
	Organization model.Organization `json:"organization"`
	Role         model.OrgRole      `json:"role"`
	JoinedAt     string             `json:"joined_at"`
}

type impl struct {
	db pg.ContextExecutor
}

func New(db pg.ContextExecutor) Repository {
	return impl{
		db: db,
	}
}
