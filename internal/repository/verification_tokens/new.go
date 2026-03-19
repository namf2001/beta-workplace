package verification_tokens

import (
	"github.com/namf2001/beta-workplace/internal/repository/db/pg"
)

type impl struct {
	db pg.ContextExecutor
}

// New creates a new verification_tokens repository
func New(db pg.ContextExecutor) Repository {
	return &impl{
		db: db,
	}
}
