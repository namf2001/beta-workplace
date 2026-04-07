package organizations

import (
	"context"

	pkgerrors "github.com/pkg/errors"

	"github.com/namf2001/beta-workplace/internal/model"
	"github.com/namf2001/beta-workplace/internal/repository"
)

// RemoveMember removes a member from organization (admin only)
// Uses transaction to ensure data consistency
func (i impl) RemoveMember(ctx context.Context, userID, orgID, memberID int64) error {
	// Use transaction to verify permissions and remove member
	return i.repo.DoInTx(ctx, func(ctx context.Context, txRepo repository.Registry) error {
		// Verify user is an admin
		member, err := txRepo.Organization().GetMember(ctx, orgID, userID)
		if err != nil {
			return err
		}

		if member.Role != model.OrgRoleAdmin {
			return pkgerrors.WithStack(ErrUnauthorized)
		}

		// Verify target member exists
		targetMember, err := txRepo.Organization().GetMember(ctx, orgID, memberID)
		if err != nil {
			return err
		}

		// Prevent removing the last admin
		if targetMember.Role == model.OrgRoleAdmin {
			adminCount, _, _, err := txRepo.Organization().GetMemberCounts(ctx, orgID)
			if err != nil {
				return err
			}
			if adminCount <= 1 {
				return pkgerrors.WithStack(ErrUnauthorized) // Cannot remove the last admin
			}
		}

		// Remove member
		return txRepo.Organization().RemoveMember(ctx, orgID, memberID)
	}, nil)
}
