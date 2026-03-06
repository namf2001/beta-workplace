package users

import (
	"github.com/namf2001/beta-workplace/internal/controller/users"
)

// Handler for api device
type Handler struct {
	userCtrl users.Controller
}

// New returns a new Handler
func New(userCtrl users.Controller) *Handler {
	return &Handler{
		userCtrl: userCtrl,
	}
}
