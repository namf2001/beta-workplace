package organizations

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/namf2001/beta-workplace/constants"
	ctrlOrgs "github.com/namf2001/beta-workplace/internal/controller/organizations"
	"github.com/namf2001/beta-workplace/internal/handler/response"
	"github.com/namf2001/beta-workplace/internal/pkg/validator"
)

// JoinOrganizationRequest represents the request body for joining an organization
type JoinOrganizationRequest struct {
	InviteCode string `json:"invite_code" validate:"required"`
}

// JoinOrganization handles POST /auth/organizations/join
// @Summary      Join organization
// @Description  Join an organization using an invite code
// @Tags         organizations
// @Accept       json
// @Produce      json
// @Param        input body JoinOrganizationRequest true "Invite code"
// @Success      200 {object} response.Response
// @Failure      400 {object} response.Response
// @Failure      500 {object} response.Response
// @Security     BearerAuth
// @Router       /auth/organizations/join [post]
func (h Handler) JoinOrganization() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req JoinOrganizationRequest
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

		member, err := h.orgCtrl.JoinOrganization(c.Request.Context(), uid, req.InviteCode)
		if err != nil {
			if errors.Is(err, ctrlOrgs.ErrInvitationExpired) {
				c.JSON(http.StatusBadRequest, response.NewResponse(
					constants.JoinOrganizationFail.Code,
					"Invitation has expired",
					nil,
				))
				return
			}
			if errors.Is(err, ctrlOrgs.ErrInvitationAlreadyUsed) {
				c.JSON(http.StatusBadRequest, response.NewResponse(
					constants.JoinOrganizationFail.Code,
					"Invitation has already been used",
					nil,
				))
				return
			}
			c.JSON(http.StatusInternalServerError, response.NewResponse(
				constants.JoinOrganizationFail.Code,
				constants.JoinOrganizationFail.Message,
				nil,
			))
			return
		}

		// Map controller output → handler-layer DTO
		res := JoinOrganizationResponse{
			ID:             member.ID,
			OrganizationID: member.OrganizationID,
			UserID:         member.UserID,
			UserName:       member.UserName,
			UserEmail:      member.UserEmail,
			UserAvatar:     member.UserAvatar,
			Role:           string(member.Role),
			InvitedBy:      member.InvitedBy,
			JoinedAt:       member.JoinedAt,
		}

		c.JSON(http.StatusOK, response.NewResponse(
			constants.JoinOrganizationSuccess.Code,
			constants.JoinOrganizationSuccess.Message,
			res,
		))
	}
}
