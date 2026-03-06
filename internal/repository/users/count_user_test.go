package users

import (
	"context"
	"testing"

	"github.com/namf2001/beta-workplace/internal/pkg/testdb"
	"github.com/namf2001/beta-workplace/internal/repository/db/pg"
	"github.com/stretchr/testify/require"
)

func TestCountUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		testdb.WithTx(t, func(tx pg.ContextExecutor) {
			testdb.LoadTestSQLFile(t, tx, "testdata/users.sql")
			repo := New(tx)
			count, err := repo.CountUser(context.Background())

			require.NoError(t, err)
			require.Equal(t, int64(3), count)
		})
	})
}
