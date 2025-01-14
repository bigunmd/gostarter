//go:build unit

package heroes

import (
	"context"
	"testing"
	"time"

	"github.com/bigunmd/gostarter/pkg/util/tests"
	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/require"
)

func TestInMemStore(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	ctx = tests.SetupTestLogger(t).WithContext(ctx)

	repo := NewInMem(ctx)

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
