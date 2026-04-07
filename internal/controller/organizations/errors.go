package organizations

import "errors"

var (
	ErrOrganizationNotFound = errors.New("organization not found")
	ErrUnauthorized         = errors.New("user not authorized for this action")
	ErrMemberNotFound       = errors.New("member not found")
	ErrInvalidRole          = errors.New("invalid role")
	ErrInvitationNotFound   = errors.New("invitation not found")
	ErrInvitationExpired    = errors.New("invitation has expired")
	ErrInvitationAlreadyUsed = errors.New("invitation has already been used")
)
