package organizations

import (
	"context"
	"database/sql"

	"github.com/namf2001/beta-workplace/internal/model"
	pkgerrors "github.com/pkg/errors"
)

// GetByID is a function that gets an organization by ID
func (i impl) GetByID(ctx context.Context, id int64) (model.Organization, error) {
	query := `
		SELECT id, name, slug, industry, size, region, logo_url, description, created_by, created_at, updated_at
		FROM organizations
		WHERE id = $1
	`

	var org model.Organization
	err := i.db.QueryRowContext(ctx, query, id).Scan(
		&org.ID,
		&org.Name,
		&org.Slug,
		&org.Industry,
		&org.Size,
		&org.Region,
		&org.LogoURL,
		&org.Description,
		&org.CreatedBy,
		&org.CreatedAt,
		&org.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return model.Organization{}, pkgerrors.WithStack(ErrOrganizationNotFound)
		}
		return model.Organization{}, pkgerrors.WithStack(err)
	}

	return org, nil
}

// GetByUserID is a function that gets organizations by user ID
func (i impl) GetByUserID(ctx context.Context, userID int64, page, pageSize int) ([]OrganizationWithRole, int64, error) {
	offset := (page - 1) * pageSize

	query := `
		SELECT
			o.id, o.name, o.slug, o.industry, o.size, o.region, o.logo_url, o.description, o.created_by, o.created_at, o.updated_at,
			om.role,
			om.joined_at
		FROM organizations o
		INNER JOIN organization_members om ON o.id = om.organization_id
		WHERE om.user_id = $1
		ORDER BY om.joined_at DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := i.db.QueryContext(ctx, query, userID, pageSize, offset)
	if err != nil {
		return nil, 0, pkgerrors.WithStack(err)
	}
	defer rows.Close()

	var orgs []OrganizationWithRole
	for rows.Next() {
		var org OrganizationWithRole
		err := rows.Scan(
			&org.Organization.ID,
			&org.Organization.Name,
			&org.Organization.Slug,
			&org.Organization.Industry,
			&org.Organization.Size,
			&org.Organization.Region,
			&org.Organization.LogoURL,
			&org.Organization.Description,
			&org.Organization.CreatedBy,
			&org.Organization.CreatedAt,
			&org.Organization.UpdatedAt,
			&org.Role,
			&org.JoinedAt,
		)
		if err != nil {
			return nil, 0, pkgerrors.WithStack(err)
		}
		orgs = append(orgs, org)
	}

	// Get total count
	countQuery := `
		SELECT COUNT(*)
		FROM organization_members
		WHERE user_id = $1
	`
	var total int64
	err = i.db.QueryRowContext(ctx, countQuery, userID).Scan(&total)
	if err != nil {
		return nil, 0, pkgerrors.WithStack(err)
	}

	return orgs, total, nil
}

func (i impl) GetMember(ctx context.Context, orgID, userID int64) (model.OrganizationMember, error) {
	query := `
		SELECT om.id, om.organization_id, om.user_id, om.role, om.invited_by, om.joined_at,
		       u.full_name, u.email, COALESCE(u.profile_image, '')
		FROM organization_members om
		INNER JOIN users u ON u.id = om.user_id
		WHERE om.organization_id = $1 AND om.user_id = $2
	`

	var member model.OrganizationMember
	err := i.db.QueryRowContext(ctx, query, orgID, userID).Scan(
		&member.ID,
		&member.OrganizationID,
		&member.UserID,
		&member.Role,
		&member.InvitedBy,
		&member.JoinedAt,
		&member.UserName,
		&member.UserEmail,
		&member.UserAvatar,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return model.OrganizationMember{}, pkgerrors.WithStack(ErrMemberNotFound)
		}
		return model.OrganizationMember{}, pkgerrors.WithStack(err)
	}

	return member, nil
}

// GetMembers is a function that gets members of an organization
func (i impl) GetMembers(ctx context.Context, orgID int64) ([]model.OrganizationMember, error) {
	query := `
		SELECT om.id, om.organization_id, om.user_id, om.role, om.invited_by, om.joined_at,
		       u.full_name, u.email, COALESCE(u.profile_image, '')
		FROM organization_members om
		INNER JOIN users u ON u.id = om.user_id
		WHERE om.organization_id = $1
		ORDER BY om.joined_at ASC
	`

	rows, err := i.db.QueryContext(ctx, query, orgID)
	if err != nil {
		return nil, pkgerrors.WithStack(err)
	}
	defer rows.Close()

	var members []model.OrganizationMember
	for rows.Next() {
		var member model.OrganizationMember
		err := rows.Scan(
			&member.ID,
			&member.OrganizationID,
			&member.UserID,
			&member.Role,
			&member.InvitedBy,
			&member.JoinedAt,
			&member.UserName,
			&member.UserEmail,
			&member.UserAvatar,
		)
		if err != nil {
			return nil, pkgerrors.WithStack(err)
		}
		members = append(members, member)
	}

	return members, nil
}

// GetMemberCounts is a function that gets member counts of an organization
func (i impl) GetMemberCounts(ctx context.Context, orgID int64) (admin, subAdmin, member int, err error) {
	query := `
		SELECT
			COALESCE(SUM(CASE WHEN role = 'admin' THEN 1 ELSE 0 END), 0) as admin_count,
			COALESCE(SUM(CASE WHEN role = 'sub_admin' THEN 1 ELSE 0 END), 0) as sub_admin_count,
			COALESCE(SUM(CASE WHEN role = 'member' THEN 1 ELSE 0 END), 0) as member_count
		FROM organization_members
		WHERE organization_id = $1
	`

	err = i.db.QueryRowContext(ctx, query, orgID).Scan(&admin, &subAdmin, &member)
	if err != nil {
		return 0, 0, 0, pkgerrors.WithStack(err)
	}

	return admin, subAdmin, member, nil
}
