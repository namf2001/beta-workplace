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

// UpdateOrganizationRequest represents the request body for updating an organization
type UpdateOrganizationRequest struct {
	Name        string `json:"name"`
	Industry    string `json:"industry"`
	Size        int    `json:"size"`
	Region      string `json:"region"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
}

// UpdateOrganization handles PUT /auth/organizations/:id
// @Summary      Update organization
// @Description  Update an organization (admin only)
// @Tags         organizations
// @Accept       json
// @Produce      json
// @Param        id    path int true "Organization ID"
// @Param        input body UpdateOrganizationRequest true "Updated fields"
// @Success      200 {object} response.Response
// @Failure      400 {object} response.Response
// @Failure      403 {object} response.Response
// @Failure      404 {object} response.Response
// @Failure      500 {object} response.Response
// @Security     BearerAuth
// @Router       /auth/organizations/{id} [put]
func (h Handler) UpdateOrganization() gin.HandlerFunc {
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

		var req UpdateOrganizationRequest
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

		input := ctrlOrgs.UpdateOrganizationInput{
			Name:        req.Name,
			Industry:    req.Industry,
			Size:        req.Size,
			Region:      req.Region,
			Avatar:      req.Avatar,
			Description: req.Description,
		}

		result, err := h.orgCtrl.UpdateOrganization(c.Request.Context(), uid, orgID, input)
		if err != nil {
			if errors.Is(err, ctrlOrgs.ErrUnauthorized) {
				c.JSON(http.StatusForbidden, response.NewResponse(
					constants.OnlyAdminCanUpdate.Code,
					constants.OnlyAdminCanUpdate.Message,
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
				constants.UpdateOrganizationFail.Code,
				constants.UpdateOrganizationFail.Message,
				nil,
			))
			return
		}

		c.JSON(http.StatusOK, response.NewResponse(
			constants.UpdateOrganizationSuccess.Code,
			constants.UpdateOrganizationSuccess.Message,
			result,
		))
	}
}
