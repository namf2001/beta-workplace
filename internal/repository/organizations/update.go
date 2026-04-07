package organizations

import (
	"context"

	"github.com/namf2001/beta-workplace/internal/model"
	pkgerrors "github.com/pkg/errors"
)

// Update updates an existing organization
func (i impl) Update(ctx context.Context, org model.Organization) error {
	org.Prepare()

	query := `
		UPDATE organizations
		SET name = $1, slug = $2, industry = $3, size = $4, region = $5, logo_url = $6, description = $7, updated_at = $8
		WHERE id = $9
	`

	result, err := i.db.ExecContext(ctx, query,
		org.Name,
		org.Slug,
		org.Industry,
		org.Size,
		org.Region,
		org.LogoURL,
		org.Description,
		org.UpdatedAt,
		org.ID,
	)

	if err != nil {
		return pkgerrors.WithStack(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return pkgerrors.WithStack(err)
	}

	if rowsAffected == 0 {
		return pkgerrors.WithStack(ErrOrganizationNotFound)
	}

	return nil
}

// UpdateMemberRole is a function that updates the role of a member in an organization
func (i impl) UpdateMemberRole(ctx context.Context, orgID, userID int64, role model.OrgRole) error {
	query := `
		UPDATE organization_members
		SET role = $1
		WHERE organization_id = $2 AND user_id = $3
	`

	result, err := i.db.ExecContext(ctx, query, role, orgID, userID)
	if err != nil {
		return pkgerrors.WithStack(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return pkgerrors.WithStack(err)
	}

	if rowsAffected == 0 {
		return pkgerrors.WithStack(ErrMemberNotFound)
	}

	return nil
}
