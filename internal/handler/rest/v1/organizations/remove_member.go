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

// RemoveMember handles DELETE /auth/organizations/:id/members/:memberId
// @Summary      Remove member
// @Description  Remove a member from the organization (admin only)
// @Tags         organizations
// @Produce      json
// @Param        id       path int true "Organization ID"
// @Param        memberId path int true "Member user ID"
// @Success      200 {object} response.Response
// @Failure      403 {object} response.Response
// @Failure      404 {object} response.Response
// @Failure      500 {object} response.Response
// @Security     BearerAuth
// @Router       /auth/organizations/{id}/members/{memberId} [delete]
func (h Handler) RemoveMember() gin.HandlerFunc {
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

		memberID, err := strconv.ParseInt(c.Param("memberId"), 10, 64)
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

		err = h.orgCtrl.RemoveMember(c.Request.Context(), uid, orgID, memberID)
		if err != nil {
			if errors.Is(err, ctrlOrgs.ErrUnauthorized) {
				c.JSON(http.StatusForbidden, response.NewResponse(
					constants.OnlyAdminCanRemoveMembers.Code,
					constants.OnlyAdminCanRemoveMembers.Message,
					nil,
				))
				return
			}
			if errors.Is(err, repoOrgs.ErrMemberNotFound) || errors.Is(err, ctrlOrgs.ErrMemberNotFound) {
				c.JSON(http.StatusNotFound, response.NewResponse(
					constants.RemoveMemberFail.Code,
					"Member not found",
					nil,
				))
				return
			}
			c.JSON(http.StatusInternalServerError, response.NewResponse(
				constants.RemoveMemberFail.Code,
				constants.RemoveMemberFail.Message,
				nil,
			))
			return
		}

		c.JSON(http.StatusOK, response.NewResponse(
			constants.RemoveMemberSuccess.Code,
			constants.RemoveMemberSuccess.Message,
			nil,
		))
	}
}
