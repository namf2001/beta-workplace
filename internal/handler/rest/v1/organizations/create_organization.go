package organizations

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/namf2001/beta-workplace/constants"
	ctrlOrgs "github.com/namf2001/beta-workplace/internal/controller/organizations"
	"github.com/namf2001/beta-workplace/internal/handler/response"
	"github.com/namf2001/beta-workplace/internal/pkg/validator"
)

// CreateOrganizationRequest represents the request body for creating an organization
type CreateOrganizationRequest struct {
	Name        string `json:"name"        validate:"required,min=1,max=255"`
	Industry    string `json:"industry"`
	Region      string `json:"region"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
	Size        int    `json:"size"`
}

// CreateOrganization handles POST /auth/organizations
// @Summary      Create organization
// @Description  Create a new organization with the current user as admin
// @Tags         organizations
// @Accept       json
// @Produce      json
// @Param        input body CreateOrganizationRequest true "Organization info"
// @Success      201  {object} response.Response
// @Failure      400  {object} response.Response
// @Failure      500  {object} response.Response
// @Security     BearerAuth
// @Router       /auth/organizations [post]
func (h Handler) CreateOrganization() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateOrganizationRequest
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

		input := ctrlOrgs.CreateOrganizationInput{
			Name:        req.Name,
			Industry:    req.Industry,
			Region:      req.Region,
			Avatar:      req.Avatar,
			Description: req.Description,
			Size:        req.Size,
		}

		org, err := h.orgCtrl.CreateOrganization(c.Request.Context(), uid, input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.NewResponse(
				constants.CreateOrganizationFail.Code,
				err.Error(),
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

		c.JSON(http.StatusCreated, response.NewResponse(
			constants.CreateOrganizationSuccess.Code,
			constants.CreateOrganizationSuccess.Message,
			res,
		))
	}
}

// GetUserOrganizations handles GET /auth/organizations
// @Summary      Get user organizations
// @Description  Get all organizations the current user is a member of
// @Tags         organizations
// @Produce      json
// @Param        page      query int true "Page number"
// @Param        page_size query int true "Page size"
// @Success      200 {object} response.Response
// @Failure      500 {object} response.Response
// @Security     BearerAuth
// @Router       /auth/organizations [get]
func (h Handler) GetUserOrganizations() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := c.Get("userID")
		uid, _ := userID.(int64)

		page := parseIntQuery(c, "page", 1)
		pageSize := parseIntQuery(c, "page_size", 20)
		if page < 1 {
			page = 1
		}
		if pageSize < 1 || pageSize > 100 {
			pageSize = 20
		}

		orgs, err := h.orgCtrl.GetUserOrganizations(c.Request.Context(), uid, page, pageSize)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.NewResponse(
				constants.GetUserOrganizationsFail.Code,
				constants.GetUserOrganizationsFail.Message,
				nil,
			))
			return
		}

		// Map controller output → handler-layer DTO
		orgList := make([]OrganizationWithRoleResponse, 0, len(orgs.Organizations))
		for _, o := range orgs.Organizations {
			orgList = append(orgList, OrganizationWithRoleResponse{
				Organization: OrganizationSummaryResponse{
					ID:       o.Organization.ID,
					Name:     o.Organization.Name,
					Industry: o.Organization.Industry,
				},
				Role:     string(o.Role),
				JoinedAt: o.JoinedAt,
			})
		}

		res := UserOrganizationsResponse{
			Organizations: orgList,
			Total:         orgs.Total,
		}

		c.JSON(http.StatusOK, response.NewResponse(
			constants.GetUserOrganizationsSuccess.Code,
			constants.GetUserOrganizationsSuccess.Message,
			res,
		))
	}
}

// parseIntQuery parses an integer query parameter, returning defaultVal on failure
func parseIntQuery(c *gin.Context, key string, defaultVal int) int {
	val, err := strconv.Atoi(c.Query(key))
	if err != nil {
		return defaultVal
	}
	return val
}

