// Package tests provides repeated utility functions.
package tests

import (
	"context"
	"testing"

	"github.com/bigunmd/gostarter/pkg/util/postgres"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

// SetupTestPostgresConfig returns [postgres.PostgresConfig] formed via
// testing environment variables.
func SetupTestPostgresConfig(t *testing.T) postgres.PostgresConfig {
	pgCfg := struct {
		Postgres postgres.PostgresConfig `env-prefix:"POSTGRES_"`
	}{}
	err := cleanenv.ReadEnv(&pgCfg)
	require.NoError(t, err)

	return pgCfg.Postgres
}

// SetupTestPostgresConn returns pgx connection to the test Postgres database.
// Additional connection options can be passed in a form of "key=value" pairs.
func SetupTestPostgresConn(ctx context.Context, t *testing.T, opts ...string) *pgx.Conn {
	pgCfg := SetupTestPostgresConfig(t)
	connCfg, err := pgx.ParseConfig(pgCfg.String(opts...))
	require.NoError(t, err)
	conn, err := pgx.ConnectConfig(ctx, connCfg)
	require.NoError(t, err)

	return conn
}

// SetupTestPostgresPool returns pgx connection pool to the test Postgres database.
// Additional connection options can be passed in a form of "key=value" pairs.
func SetupTestPostgresPool(ctx context.Context, t *testing.T, opts ...string) *pgxpool.Pool {
	pgCfg := SetupTestPostgresConfig(t)
	poolCfg, err := pgxpool.ParseConfig(pgCfg.String(opts...))
	require.NoError(t, err)
	pool, err := pgxpool.NewWithConfig(ctx, poolCfg)
	require.NoError(t, err)

	return pool
}
