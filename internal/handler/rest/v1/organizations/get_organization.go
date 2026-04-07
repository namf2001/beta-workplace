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

// GetOrganization handles GET /auth/organizations/:id
// @Summary      Get organization details
// @Description  Get details of a specific organization by ID
// @Tags         organizations
// @Produce      json
// @Param        id path int true "Organization ID"
// @Success      200 {object} response.Response
// @Failure      403 {object} response.Response
// @Failure      404 {object} response.Response
// @Failure      500 {object} response.Response
// @Security     BearerAuth
// @Router       /auth/organizations/{id} [get]
func (h Handler) GetOrganization() gin.HandlerFunc {
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

		org, err := h.orgCtrl.GetOrganization(c.Request.Context(), uid, orgID)
		if err != nil {
			if errors.Is(err, ctrlOrgs.ErrUnauthorized) {
				c.JSON(http.StatusForbidden, response.NewResponse(
					constants.AccessDeniedToOrganization.Code,
					constants.AccessDeniedToOrganization.Message,
					nil,
				))
				return
			}
			if errors.Is(err, repoOrgs.ErrOrganizationNotFound) {
				c.JSON(http.StatusNotFound, response.NewResponse(
					constants.OrganizationNotFound.Code,
					constants.OrganizationNotFound.Message,
					nil,
				))
				return
			}
			c.JSON(http.StatusInternalServerError, response.NewResponse(
				constants.InternalServerError.Code,
				constants.InternalServerError.Message,
				nil,
			))
			return
		}

		// Map controller output → handler-layer DTO
		res := OrganizationDetailResponse{
			ID:            org.ID,
			Name:          org.Name,
			Industry:      org.Industry,
			Size:          org.Size,
			Region:        org.Region,
			Avatar:        org.Avatar,
			Description:   org.Description,
			AdminCount:    org.AdminCount,
			MemberCount:   org.MemberCount,
			SubAdminCount: org.SubAdminCount,
			CreatedAt:     org.CreatedAt,
			UpdatedAt:     org.UpdatedAt,
		}

		c.JSON(http.StatusOK, response.NewResponse(
			constants.GetOrganizationSuccess.Code,
			constants.GetOrganizationSuccess.Message,
			res,
		))
	}
}
