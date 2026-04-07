package organizations

import (
	"context"
)

// GetOrganization gets organization details by ID
func (i impl) GetOrganization(ctx context.Context, userID, orgID int64) (OrganizationResponse, error) {
	// Verify user is a member of the organization
	_, err := i.repo.Organization().GetMember(ctx, orgID, userID)
	if err != nil {
		return OrganizationResponse{}, err
	}

	org, err := i.repo.Organization().GetByID(ctx, orgID)
	if err != nil {
		return OrganizationResponse{}, err
	}

	adminCount, subAdminCount, memberCount, err := i.repo.Organization().GetMemberCounts(ctx, orgID)
	if err != nil {
		return OrganizationResponse{}, err
	}

	return OrganizationResponse{
		ID:            org.ID,
		Name:          org.Name,
		Industry:      org.Industry,
		Size:          org.Size,
		Region:        org.Region,
		Avatar:        org.LogoURL,
		Description:   org.Description,
		AdminCount:    adminCount,
		MemberCount:   adminCount + subAdminCount + memberCount,
		SubAdminCount: subAdminCount,
		CreatedAt:     org.CreatedAt,
		UpdatedAt:     org.UpdatedAt,
	}, nil
}

// GetUserOrganizations gets all organizations where user is a member
func (i impl) GetUserOrganizations(ctx context.Context, userID int64, page, pageSize int) (OrganizationListResponse, error) {
	orgs, total, err := i.repo.Organization().GetByUserID(ctx, userID, page, pageSize)
	if err != nil {
		return OrganizationListResponse{}, err
	}

	organizations := make([]OrganizationWithRoleResponse, 0, len(orgs))
	for _, org := range orgs {
		organizations = append(organizations, OrganizationWithRoleResponse{
			Organization: OrganizationSummary{
				ID:       org.Organization.ID,
				Name:     org.Organization.Name,
				Industry: org.Organization.Industry,
			},
			Role:     org.Role,
			JoinedAt: org.JoinedAt,
		})
	}

	return OrganizationListResponse{
		Organizations: organizations,
		Total:         total,
	}, nil
}

// GetMembers gets all members of an organization
func (i impl) GetMembers(ctx context.Context, userID, orgID int64) ([]MemberResponse, error) {
	// Verify user is a member of the organization
	_, err := i.repo.Organization().GetMember(ctx, orgID, userID)
	if err != nil {
		return nil, err
	}

	members, err := i.repo.Organization().GetMembers(ctx, orgID)
	if err != nil {
		return nil, err
	}

	result := make([]MemberResponse, 0, len(members))
	for _, m := range members {
		result = append(result, MemberResponse{
			ID:             m.ID,
			OrganizationID: m.OrganizationID,
			UserID:         m.UserID,
			UserName:       m.UserName,
			UserEmail:      m.UserEmail,
			UserAvatar:     m.UserAvatar,
			Role:           m.Role,
			InvitedBy:      m.InvitedBy,
			JoinedAt:       m.JoinedAt,
		})
	}

	// ensure we return empty slice not nil
	if result == nil {
		result = []MemberResponse{}
	}

	return result, nil
}
