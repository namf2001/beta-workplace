package organizations

import (
	"context"

	pkgerrors "github.com/pkg/errors"

	"github.com/namf2001/beta-workplace/internal/model"
	"github.com/namf2001/beta-workplace/internal/repository"
)

// UpdateOrganization updates organization information (admin only)
func (i impl) UpdateOrganization(ctx context.Context, userID, orgID int64, input UpdateOrganizationInput) (UpdateOrganizationResponse, error) {
	var updatedOrg model.Organization

	err := i.repo.DoInTx(ctx, func(ctx context.Context, txRepo repository.Registry) error {
		// Verify user is an admin
		member, err := txRepo.Organization().GetMember(ctx, orgID, userID)
		if err != nil {
			return err
		}

		if member.Role != model.OrgRoleAdmin {
			return pkgerrors.WithStack(ErrUnauthorized)
		}

		// Get existing organization
		org, err := txRepo.Organization().GetByID(ctx, orgID)
		if err != nil {
			return err
		}

		// Update organization fields (only set if provided)
		if input.Name != "" {
			org.Name = input.Name
			org.Slug = generateSlug(input.Name)
		}
		if input.Industry != "" {
			org.Industry = input.Industry
		}
		if input.Size > 0 {
			org.Size = input.Size
		}
		if input.Region != "" {
			org.Region = input.Region
		}
		if input.Avatar != "" {
			org.LogoURL = input.Avatar
		}
		if input.Description != "" {
			org.Description = input.Description
		}

		err = txRepo.Organization().Update(ctx, org)
		if err != nil {
			return err
		}

		updatedOrg = org
		return nil
	}, nil)

	if err != nil {
		return UpdateOrganizationResponse{}, err
	}

	return UpdateOrganizationResponse{
		ID:   updatedOrg.ID,
		Name: updatedOrg.Name,
	}, nil
}

// UpdateMemberRole updates member role (admin only)
func (i impl) UpdateMemberRole(ctx context.Context, userID, orgID int64, input UpdateMemberRoleInput) error {
	return i.repo.DoInTx(ctx, func(ctx context.Context, txRepo repository.Registry) error {
		member, err := txRepo.Organization().GetMember(ctx, orgID, userID)
		if err != nil {
			return err
		}

		if member.Role != model.OrgRoleAdmin {
			return pkgerrors.WithStack(ErrUnauthorized)
		}

		var role model.OrgRole
		switch input.Role {
		case "admin":
			role = model.OrgRoleAdmin
		case "sub_admin":
			role = model.OrgRoleSubAdmin
		case "member":
			role = model.OrgRoleMember
		default:
			return ErrInvalidRole
		}

		return txRepo.Organization().UpdateMemberRole(ctx, orgID, input.UserID, role)
	}, nil)
}
