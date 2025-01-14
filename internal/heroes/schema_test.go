//go:build integration

package heroes

import (
	"context"
	"testing"

	"github.com/bigunmd/gostarter/pkg/util/tests"
	"github.com/stretchr/testify/require"
)

func TestCreateSchema(t *testing.T) {
	ctx := context.Background()
	ctx = tests.SetupTestLogger(t).WithContext(ctx)

	conn := tests.SetupTestPostgresConn(ctx, t)
	t.Cleanup(func() {
		err := conn.Close(ctx)
		require.NoError(t, err)
	})

	err := createSchema(ctx, conn, "test_schema")
	require.NoError(t, err)
}

func TestDropSchema(t *testing.T) {
	ctx := context.Background()
	ctx = tests.SetupTestLogger(t).WithContext(ctx)

	conn := tests.SetupTestPostgresConn(ctx, t)
	t.Cleanup(func() {
		err := conn.Close(ctx)
		require.NoError(t, err)
	})

	err := dropSchema(ctx, conn, "test_schema")
	require.NoError(t, err)
}
