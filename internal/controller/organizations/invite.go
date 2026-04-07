package organizations

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	pkgerrors "github.com/pkg/errors"

	"github.com/namf2001/beta-workplace/internal/model"
	"github.com/namf2001/beta-workplace/internal/repository"
)

// inviteCodeLength is the number of random bytes in each invite code (= 9 chars in hex)
const inviteCodeBytes = 9
const invitationExpiryDays = 7

// generateInviteCode generates a random, URL-safe invite code
func generateInviteCode() (string, error) {
	b := make([]byte, inviteCodeBytes)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// InviteMember invites a member to organization (admin or sub-admin only)
func (i impl) InviteMember(ctx context.Context, userID, orgID int64, input InviteMemberInput) (InvitationResponse, error) {
	// Verify user is admin or sub-admin
	member, err := i.repo.Organization().GetMember(ctx, orgID, userID)
	if err != nil {
		return InvitationResponse{}, err
	}

	if member.Role != model.OrgRoleAdmin && member.Role != model.OrgRoleSubAdmin {
		return InvitationResponse{}, pkgerrors.WithStack(ErrUnauthorized)
	}

	// Validate role
	var role model.OrgRole
	switch input.Role {
	case "admin":
		role = model.OrgRoleAdmin
	case "sub_admin":
		role = model.OrgRoleSubAdmin
	case "member", "":
		role = model.OrgRoleMember
	default:
		return InvitationResponse{}, ErrInvalidRole
	}

	// Generate unique invite code
	inviteCode, err := generateInviteCode()
	if err != nil {
		return InvitationResponse{}, fmt.Errorf("generate invite code: %w", err)
	}

	inv := model.OrganizationInvitation{
		OrganizationID: orgID,
		InviteCode:     inviteCode,
		Email:          input.Email,
		Role:           role,
		InvitedBy:      userID,
		ExpiresAt:      time.Now().Add(invitationExpiryDays * 24 * time.Hour),
	}

	created, err := i.repo.Organization().CreateInvitation(ctx, inv)
	if err != nil {
		return InvitationResponse{}, err
	}

	return InvitationResponse{
		ID:             created.ID,
		OrganizationID: created.OrganizationID,
		InviteCode:     created.InviteCode,
		Email:          created.Email,
		ExpiresAt:      created.ExpiresAt,
		Status:         created.Status,
		CreatedAt:      created.CreatedAt,
	}, nil
}

// JoinOrganization joins an organization using invite code
func (i impl) JoinOrganization(ctx context.Context, userID int64, inviteCode string) (JoinOrganizationResponse, error) {
	var joinedMember model.OrganizationMember

	err := i.repo.DoInTx(ctx, func(ctx context.Context, txRepo repository.Registry) error {
		// Validate invite code
		inv, err := txRepo.Organization().GetInvitationByCode(ctx, inviteCode)
		if err != nil {
			return err
		}

		// Check invitation status
		if inv.Status != "pending" {
			return pkgerrors.WithStack(ErrInvitationAlreadyUsed)
		}

		// Check expiration
		if time.Now().After(inv.ExpiresAt) {
			return pkgerrors.WithStack(ErrInvitationExpired)
		}

		// Add user as member
		invitedBy := inv.InvitedBy
		member := model.OrganizationMember{
			OrganizationID: inv.OrganizationID,
			UserID:         userID,
			Role:           inv.Role,
			InvitedBy:      &invitedBy,
		}

		joinedMember, err = txRepo.Organization().AddMember(ctx, member)
		if err != nil {
			return err
		}

		// Mark invitation as accepted
		return txRepo.Organization().UseInvitation(ctx, inviteCode)
	}, nil)

	if err != nil {
		return JoinOrganizationResponse{}, err
	}

	// Fetch user info to return in response
	fullMember, err := i.repo.Organization().GetMember(ctx, joinedMember.OrganizationID, userID)
	if err != nil {
		// return partial info if user join succeeded but user fetch failed
		return JoinOrganizationResponse{
			ID:             joinedMember.ID,
			OrganizationID: joinedMember.OrganizationID,
			UserID:         joinedMember.UserID,
			Role:           joinedMember.Role,
			InvitedBy:      joinedMember.InvitedBy,
			JoinedAt:       joinedMember.JoinedAt,
		}, nil
	}

	return JoinOrganizationResponse{
		ID:             fullMember.ID,
		OrganizationID: fullMember.OrganizationID,
		UserID:         fullMember.UserID,
		UserName:       fullMember.UserName,
		UserEmail:      fullMember.UserEmail,
		UserAvatar:     fullMember.UserAvatar,
		Role:           fullMember.Role,
		InvitedBy:      fullMember.InvitedBy,
		JoinedAt:       fullMember.JoinedAt,
	}, nil
}
