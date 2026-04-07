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
)

// UpdateMemberRoleRequest represents the request body for updating a member's role
type UpdateMemberRoleRequest struct {
	UserID int64  `json:"user_id" validate:"required"`
	Role   string `json:"role"    validate:"required,oneof=admin sub_admin member"`
}

// UpdateMemberRole handles PUT /auth/organizations/:id/members/role
// @Summary      Update member role
// @Description  Update the role of a member in the organization (admin only)
// @Tags         organizations
// @Accept       json
// @Produce      json
// @Param        id    path int true "Organization ID"
// @Param        input body UpdateMemberRoleRequest true "Role update"
// @Success      200 {object} response.Response
// @Failure      400 {object} response.Response
// @Failure      403 {object} response.Response
// @Failure      500 {object} response.Response
// @Security     BearerAuth
// @Router       /auth/organizations/{id}/members/role [put]
func (h Handler) UpdateMemberRole() gin.HandlerFunc {
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

		var req UpdateMemberRoleRequest
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

		input := ctrlOrgs.UpdateMemberRoleInput{
			UserID: req.UserID,
			Role:   req.Role,
		}

		err = h.orgCtrl.UpdateMemberRole(c.Request.Context(), uid, orgID, input)
		if err != nil {
			if errors.Is(err, ctrlOrgs.ErrUnauthorized) {
				c.JSON(http.StatusForbidden, response.NewResponse(
					constants.OnlyAdminCanUpdateMemberRole.Code,
					constants.OnlyAdminCanUpdateMemberRole.Message,
					nil,
				))
				return
			}
			if errors.Is(err, ctrlOrgs.ErrInvalidRole) {
				c.JSON(http.StatusBadRequest, response.NewResponse(
					constants.InvalidRole.Code,
					constants.InvalidRole.Message,
					nil,
				))
				return
			}
			c.JSON(http.StatusInternalServerError, response.NewResponse(
				constants.UpdateMemberRoleFail.Code,
				constants.UpdateMemberRoleFail.Message,
				nil,
			))
			return
		}

		c.JSON(http.StatusOK, response.NewResponse(
			constants.UpdateMemberRoleSuccess.Code,
			constants.UpdateMemberRoleSuccess.Message,
			nil,
		))
	}
}
