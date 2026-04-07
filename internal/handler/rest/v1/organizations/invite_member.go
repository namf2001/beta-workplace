package organizations

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/namf2001/beta-workplace/constants"
	ctrlOrgs "github.com/namf2001/beta-workplace/internal/controller/organizations"
	"github.com/namf2001/beta-workplace/internal/handler/response"
	"github.com/namf2001/beta-workplace/internal/pkg/validator"
	repoOrgs "github.com/namf2001/beta-workplace/internal/repository/organizations"
)

// InviteMemberRequest represents the request body for inviting a member
type InviteMemberRequest struct {
	Email string `json:"email" validate:"required,email"`
	Role  string `json:"role"  validate:"required,oneof=admin sub_admin member"`
}

// InviteMember handles POST /auth/organizations/:id/invite
// @Summary      Invite member
// @Description  Invite a new member to the organization (admin or sub-admin only)
// @Tags         organizations
// @Accept       json
// @Produce      json
// @Param        id    path int true "Organization ID"
// @Param        input body InviteMemberRequest true "Member invitation"
// @Success      201 {object} response.Response
// @Failure      400 {object} response.Response
// @Failure      403 {object} response.Response
// @Failure      500 {object} response.Response
// @Security     BearerAuth
// @Router       /auth/organizations/{id}/invite [post]
func (h Handler) InviteMember() gin.HandlerFunc {
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

		var req InviteMemberRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, response.NewResponse(
				constants.BindJSONFail.Code,
				constants.BindJSONFail.Message,
				nil,
			))
			return
		}

		if err := validator.Validate(req); err != nil {
			c.JSON(http.StatusBadRequest, response.NewResponse(
				constants.InvalidRequestParams.Code,
				constants.InvalidRequestParams.Message,
				nil,
			))
			return
		}

		userID, _ := c.Get("userID")
		uid, _ := userID.(int64)

		input := ctrlOrgs.InviteMemberInput{
			Email: req.Email,
			Role:  req.Role,
		}

		inv, err := h.orgCtrl.InviteMember(c.Request.Context(), uid, orgID, input)
		if err != nil {
			if errors.Is(err, ctrlOrgs.ErrUnauthorized) {
				c.JSON(http.StatusForbidden, response.NewResponse(
					constants.OnlyAdminOrSubAdminCanInvite.Code,
					constants.OnlyAdminOrSubAdminCanInvite.Message,
					nil,
				))
				return
			}
			if errors.Is(err, ctrlOrgs.ErrInvalidRole) || errors.Is(err, repoOrgs.ErrMemberNotFound) {
				c.JSON(http.StatusBadRequest, response.NewResponse(
					constants.InvalidRole.Code,
					constants.InvalidRole.Message,
					nil,
				))
				return
			}
			c.JSON(http.StatusInternalServerError, response.NewResponse(
				constants.CreateInvitationFail.Code,
				constants.CreateInvitationFail.Message,
				nil,
			))
			return
		}

		// Map controller output → handler-layer DTO
		res := InvitationDetailResponse{
			ID:             inv.ID,
			OrganizationID: inv.OrganizationID,
			InviteCode:     inv.InviteCode,
			Email:          inv.Email,
			ExpiresAt:      inv.ExpiresAt,
			Status:         inv.Status,
			CreatedAt:      inv.CreatedAt,
		}

		c.JSON(http.StatusCreated, response.NewResponse(
			constants.CreateInvitationSuccess.Code,
			constants.CreateInvitationSuccess.Message,
			res,
		))
	}
}
