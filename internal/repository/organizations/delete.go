package organizations

import (
	"context"

	pkgerrors "github.com/pkg/errors"
)

// Delete is a function that deletes an organization
func (i impl) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM organizations WHERE id = $1`

	result, err := i.db.ExecContext(ctx, query, id)
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

// RemoveMember is a function that removes a member from an organization
func (i impl) RemoveMember(ctx context.Context, orgID, userID int64) error {
	query := `DELETE FROM organization_members WHERE organization_id = $1 AND user_id = $2`

	result, err := i.db.ExecContext(ctx, query, orgID, userID)
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
