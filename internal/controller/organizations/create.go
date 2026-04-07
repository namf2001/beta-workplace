package organizations

import (
	"context"
	"strings"

	"github.com/namf2001/beta-workplace/internal/model"
	"github.com/namf2001/beta-workplace/internal/repository"
)

// CreateOrganization creates a new organization with the user as admin
func (i impl) CreateOrganization(ctx context.Context, userID int64, input CreateOrganizationInput) (OrganizationResponse, error) {
	var createdOrg model.Organization

	err := i.repo.DoInTx(ctx, func(ctx context.Context, txRepo repository.Registry) error {
		org := model.Organization{
			Name:        input.Name,
			Slug:        generateSlug(input.Name),
			Industry:    input.Industry,
			Size:        input.Size,
			Region:      input.Region,
			LogoURL:     input.Avatar,
			Description: input.Description,
			CreatedBy:   userID,
		}

		var err error
		createdOrg, err = txRepo.Organization().Create(ctx, org)
		if err != nil {
			return err
		}

		member := model.OrganizationMember{
			OrganizationID: createdOrg.ID,
			UserID:         userID,
			Role:           model.OrgRoleAdmin,
		}

		_, err = txRepo.Organization().AddMember(ctx, member)
		if err != nil {
			return err
		}

		return nil
	}, nil)

	if err != nil {
		return OrganizationResponse{}, err
	}

	adminCount, subAdminCount, memberCount, err := i.repo.Organization().GetMemberCounts(ctx, createdOrg.ID)
	if err != nil {
		return OrganizationResponse{}, err
	}

	return OrganizationResponse{
		ID:            createdOrg.ID,
		Name:          createdOrg.Name,
		Industry:      createdOrg.Industry,
		Size:          createdOrg.Size,
		Region:        createdOrg.Region,
		Avatar:        createdOrg.LogoURL,
		Description:   createdOrg.Description,
		AdminCount:    adminCount,
		MemberCount:   adminCount + subAdminCount + memberCount,
		SubAdminCount: subAdminCount,
		CreatedAt:     createdOrg.CreatedAt,
		UpdatedAt:     createdOrg.UpdatedAt,
	}, nil
}

// generateSlug generates a URL-friendly slug from organization name
func generateSlug(name string) string {
	slug := strings.ToLower(name)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			return r
		}
		return -1
	}, slug)
	return slug
}
