package organizations

import (
	ctrlOrgs "github.com/namf2001/beta-workplace/internal/controller/organizations"
)

// Handler handles organization HTTP requests
type Handler struct {
	orgCtrl ctrlOrgs.Controller
}

// New returns a new organizations Handler
func New(orgCtrl ctrlOrgs.Controller) *Handler {
	return &Handler{
		orgCtrl: orgCtrl,
	}
}
