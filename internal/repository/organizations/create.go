package organizations

import (
	"context"

	"github.com/namf2001/beta-workplace/internal/model"
	pkgerrors "github.com/pkg/errors"
)

// Create is a function that creates a new organization
func (i impl) Create(ctx context.Context, org model.Organization) (model.Organization, error) {
	org.Prepare()

	query := `
		INSERT INTO organizations (name, slug, industry, size, region, logo_url, description, created_by, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id
	`

	err := i.db.QueryRowContext(ctx, query,
		org.Name,
		org.Slug,
		org.Industry,
		org.Size,
		org.Region,
		org.LogoURL,
		org.Description,
		org.CreatedBy,
		org.CreatedAt,
		org.UpdatedAt,
	).Scan(&org.ID)

	if err != nil {
		return model.Organization{}, pkgerrors.WithStack(err)
	}

	return org, nil
}

func (i impl) AddMember(ctx context.Context, member model.OrganizationMember) (model.OrganizationMember, error) {
	query := `
		INSERT INTO organization_members (organization_id, user_id, role, invited_by, joined_at)
		VALUES ($1, $2, $3, $4, NOW())
		RETURNING id, joined_at
	`

	err := i.db.QueryRowContext(ctx, query,
		member.OrganizationID,
		member.UserID,
		member.Role,
		member.InvitedBy,
	).Scan(&member.ID, &member.JoinedAt)

	if err != nil {
		return model.OrganizationMember{}, pkgerrors.WithStack(err)
	}

	return member, nil
}
