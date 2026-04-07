package organizations

import (
	"context"
	"database/sql"

	"github.com/namf2001/beta-workplace/internal/model"
	pkgerrors "github.com/pkg/errors"
)

// CreateInvitation creates a new organization invitation
func (i impl) CreateInvitation(ctx context.Context, inv model.OrganizationInvitation) (model.OrganizationInvitation, error) {
	query := `
		INSERT INTO organization_invitations (organization_id, invite_code, email, role, invited_by, status, expires_at, created_at)
		VALUES ($1, $2, $3, $4, $5, 'pending', $6, NOW())
		RETURNING id, created_at
	`

	err := i.db.QueryRowContext(ctx, query,
		inv.OrganizationID,
		inv.InviteCode,
		inv.Email,
		inv.Role,
		inv.InvitedBy,
		inv.ExpiresAt,
	).Scan(&inv.ID, &inv.CreatedAt)

	if err != nil {
		return model.OrganizationInvitation{}, pkgerrors.WithStack(err)
	}

	inv.Status = "pending"
	return inv, nil
}

// GetInvitationByCode retrieves an invitation by invite code
func (i impl) GetInvitationByCode(ctx context.Context, inviteCode string) (model.OrganizationInvitation, error) {
	query := `
		SELECT id, organization_id, invite_code, email, role, invited_by, status, expires_at, created_at
		FROM organization_invitations
		WHERE invite_code = $1
	`

	var inv model.OrganizationInvitation
	err := i.db.QueryRowContext(ctx, query, inviteCode).Scan(
		&inv.ID,
		&inv.OrganizationID,
		&inv.InviteCode,
		&inv.Email,
		&inv.Role,
		&inv.InvitedBy,
		&inv.Status,
		&inv.ExpiresAt,
		&inv.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return model.OrganizationInvitation{}, pkgerrors.WithStack(ErrInvitationNotFound)
		}
		return model.OrganizationInvitation{}, pkgerrors.WithStack(err)
	}

	return inv, nil
}

// UseInvitation marks an invitation as accepted
func (i impl) UseInvitation(ctx context.Context, inviteCode string) error {
	query := `
		UPDATE organization_invitations
		SET status = 'accepted'
		WHERE invite_code = $1
	`

	result, err := i.db.ExecContext(ctx, query, inviteCode)
	if err != nil {
		return pkgerrors.WithStack(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return pkgerrors.WithStack(err)
	}

	if rowsAffected == 0 {
		return pkgerrors.WithStack(ErrInvitationNotFound)
	}

	return nil
}
