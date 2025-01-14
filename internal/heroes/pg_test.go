//go:build integration

package heroes

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/bigunmd/gostarter/pkg/util/tests"
	"github.com/jaswdr/faker/v2"
	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupDB(ctx context.Context, t *testing.T, schemaName string) func() {
	conn := tests.SetupTestPostgresConn(ctx, t)
	err := createSchema(ctx, conn, schemaName)
	require.NoError(t, err)
	pgCfg := tests.SetupTestPostgresConfig(t)
	err = migrateUp(ctx, pgCfg.URL(
		fmt.Sprintf("x-migrations-table=%s_migrations", schemaName),
		fmt.Sprintf("search_path=%s", schemaName),
	))
	require.NoError(t, err)

	return func() {
		err := dropSchema(ctx, conn, schemaName)
		assert.NoError(t, err)
		err = dropSchema(ctx, conn, schemaName+"_migrations")
		assert.NoError(t, err)
		err = conn.Close(ctx)
		assert.NoError(t, err)
	}
}

func setupTestPg(ctx context.Context, t *testing.T, schemaName string) (*pg, func()) {
	pool := tests.SetupTestPostgresPool(ctx, t, "search_path="+schemaName)
	pg := NewPg(ctx, pool)

	cleanup := func() {
		pool.Close()
	}

	return pg, cleanup
}

func TestPgStore(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	ctx = tests.SetupTestLogger(t).WithContext(ctx)

	schemaName := xid.New().String()
	cleanupDB := setupDB(ctx, t, schemaName)
	t.Cleanup(cleanupDB)

	repo, cleanupPg := setupTestPg(ctx, t, schemaName)
	t.Cleanup(cleanupPg)

	f := faker.New()

	h := &Hero{
		CreatedAt: f.Time().Time(time.Now()),
		UpdatedAt: f.Time().Time(time.Now()),
		ID:        f.UUID().V4(),
		Name:      f.Person().Name(),
		Owner:     f.Person().Contact().Email,
	}

	storedHero, err := repo.Store(ctx, h)
	require.NoError(t, err)
	require.Equal(t, h, storedHero)

	_, err = repo.Store(ctx, h)
	require.ErrorIs(t, err, ErrAlreadyExists)
}
