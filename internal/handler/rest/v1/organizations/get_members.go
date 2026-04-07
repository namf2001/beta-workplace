package organizations

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/namf2001/beta-workplace/constants"
	ctrlOrgs "github.com/namf2001/beta-workplace/internal/controller/organizations"
	"github.com/namf2001/beta-workplace/internal/handler/response"
	repoOrgs "github.com/namf2001/beta-workplace/internal/repository/organizations"
)

// GetMembers handles GET /auth/organizations/:id/members
// @Summary      Get organization members
// @Description  Get all members of an organization
// @Tags         organizations
// @Produce      json
// @Param        id path int true "Organization ID"
// @Success      200 {object} response.Response
// @Failure      403 {object} response.Response
// @Failure      500 {object} response.Response
// @Security     BearerAuth
// @Router       /auth/organizations/{id}/members [get]
func (h Handler) GetMembers() gin.HandlerFunc {
	return func(c *gin.Context) {
		orgID, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.NewResponse(
				constants.InvalidRequestParams.Code,
				constants.InvalidRequestParams.Message,
				nil,
			))
			return
		}

		userID, _ := c.Get("userID")
		uid, _ := userID.(int64)

		members, err := h.orgCtrl.GetMembers(c.Request.Context(), uid, orgID)
		if err != nil {
			if errors.Is(err, ctrlOrgs.ErrUnauthorized) {
				c.JSON(http.StatusForbidden, response.NewResponse(
					constants.AccessDeniedToOrganization.Code,
					constants.AccessDeniedToOrganization.Message,
					nil,
				))
				return
			}
			if errors.Is(err, repoOrgs.ErrMemberNotFound) {
				c.JSON(http.StatusForbidden, response.NewResponse(
					constants.AccessDeniedToOrganization.Code,
					constants.AccessDeniedToOrganization.Message,
					nil,
				))
				return
			}
			c.JSON(http.StatusInternalServerError, response.NewResponse(
				constants.GetOrganizationMembersFail.Code,
				constants.GetOrganizationMembersFail.Message,
				nil,
			))
			return
		}

		// Map controller output → handler-layer DTOs
		res := make([]MemberDetailResponse, 0, len(members))
		for _, m := range members {
			res = append(res, MemberDetailResponse{
				ID:             m.ID,
				OrganizationID: m.OrganizationID,
				UserID:         m.UserID,
				UserName:       m.UserName,
				UserEmail:      m.UserEmail,
				UserAvatar:     m.UserAvatar,
				Role:           string(m.Role),
				InvitedBy:      m.InvitedBy,
				JoinedAt:       m.JoinedAt,
			})
		}

		c.JSON(http.StatusOK, response.NewResponse(
			constants.GetOrganizationMembersSuccess.Code,
			constants.GetOrganizationMembersSuccess.Message,
			res,
		))
	}
}
