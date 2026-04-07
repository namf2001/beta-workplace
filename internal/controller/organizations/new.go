package organizations

import (
	"context"

	"github.com/namf2001/beta-workplace/internal/repository"
)

type Controller interface {
	// CreateOrganization creates a new organization with the user as admin
	CreateOrganization(ctx context.Context, userID int64, input CreateOrganizationInput) (OrganizationResponse, error)

	// GetOrganization gets organization details by ID
	GetOrganization(ctx context.Context, userID, orgID int64) (OrganizationResponse, error)

	// GetUserOrganizations gets all organizations where user is a member
	GetUserOrganizations(ctx context.Context, userID int64, page, pageSize int) (OrganizationListResponse, error)

	// UpdateOrganization updates organization information (admin only)
	UpdateOrganization(ctx context.Context, userID, orgID int64, input UpdateOrganizationInput) (UpdateOrganizationResponse, error)

	// InviteMember invites a member to organization
	InviteMember(ctx context.Context, userID, orgID int64, input InviteMemberInput) (InvitationResponse, error)

	// JoinOrganization joins an organization using invite code
	JoinOrganization(ctx context.Context, userID int64, inviteCode string) (JoinOrganizationResponse, error)

	// GetMembers gets all members of an organization
	GetMembers(ctx context.Context, userID, orgID int64) ([]MemberResponse, error)

	// UpdateMemberRole updates member role (admin only)
	UpdateMemberRole(ctx context.Context, userID, orgID int64, input UpdateMemberRoleInput) error

	// RemoveMember removes a member from organization (admin only)
	RemoveMember(ctx context.Context, userID, orgID, memberID int64) error
}

type impl struct {
	repo repository.Registry
}

func New(repo repository.Registry) Controller {
	return impl{
		repo: repo,
	}
}
