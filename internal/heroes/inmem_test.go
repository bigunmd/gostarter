package heroes

import (
	"context"
	"testing"
	"time"

	"github.com/jaswdr/faker/v2"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func setupTestLogger(t *testing.T) zerolog.Logger {
	return zerolog.New(zerolog.NewTestWriter(t))
}

func TestInMemStore(t *testing.T) {
	ctx := context.Background()
	ctx = setupTestLogger(t).WithContext(ctx)

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
