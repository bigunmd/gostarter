//go:build integration

package heroes

import (
	"context"
	"fmt"
	"testing"

	"github.com/bigunmd/gostarter/pkg/util/tests"
	"github.com/jackc/pgx/v5"
	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMigrateUp(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	ctx = tests.SetupTestLogger(t).WithContext(ctx)

	pgCfg := tests.SetupTestPostgresConfig(t)

	connCfg, err := pgx.ParseConfig(pgCfg.String())
	require.NoError(t, err)

	conn, err := pgx.ConnectConfig(ctx, connCfg)
	require.NoError(t, err)
	t.Cleanup(func() {
		err := conn.Close(ctx)
		assert.NoError(t, err)
	})

	schemaName := xid.New().String()
	err = createSchema(ctx, conn, schemaName)
	require.NoError(t, err)
	t.Cleanup(func() {
		err := dropSchema(ctx, conn, schemaName)
		assert.NoError(t, err)
	})

	err = migrateUp(ctx, pgCfg.URL(
		fmt.Sprintf("x-migrations-table=%s_migrations", schemaName),
		fmt.Sprintf("search_path=%s", schemaName),
	))
	require.NoError(t, err)
}
