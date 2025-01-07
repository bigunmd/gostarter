package heroes

import (
	"context"
	"testing"
	"time"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/require"
)

func setupTestService(ctx context.Context) *service {
	repo := NewInMem(ctx)

	return NewService(ctx, repo)
}

func TestServiceCreate(t *testing.T) {
	ctx := context.Background()
	ctx = setupTestLogger(t).WithContext(ctx)

	s := setupTestService(ctx)

	f := faker.New()
	req := &CreateHeroRequest{
		Hero: Hero{
			CreatedAt: f.Time().Time(time.Now()),
			UpdatedAt: f.Time().Time(time.Now()),
			ID:        f.UUID().V4(),
			Name:      f.Person().Name(),
			Owner:     f.Person().Contact().Email,
		},
	}

	resp, err := s.Create(ctx, req)
	require.NoError(t, err)
	require.NotNil(t, resp)
}
