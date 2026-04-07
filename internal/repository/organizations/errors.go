package organizations

import "errors"

var (
	ErrOrganizationNotFound = errors.New("organization not found")
	ErrMemberNotFound       = errors.New("member not found")
	ErrInvitationNotFound   = errors.New("invitation not found")
)
