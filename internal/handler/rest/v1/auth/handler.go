package auth

import (
	"github.com/namf2001/beta-workplace/internal/controller/auth"
)

type Handler struct {
	ctrl auth.Controller
}

func New(ctrl auth.Controller) *Handler {
	return &Handler{
		ctrl: ctrl,
	}
}
