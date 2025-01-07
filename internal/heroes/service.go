package heroes

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/xid"
	"github.com/rs/zerolog"
)

// Repository declares repository interface for [Hero].
type Repository interface {
	// Store stores [Hero] in the repository.
	Store(ctx context.Context, hero *Hero) (*Hero, error)
}

var _ Service = (*service)(nil)

type service struct {
	log  *zerolog.Logger
	repo Repository
}

// Create implements [Service].
func (s *service) Create(ctx context.Context, req *CreateHeroRequest) (*CreateHeroResponse, error) {
	log := zerolog.Ctx(ctx).With().Any("req", req).Logger()

	log.Debug().Msg("creating hero")
	h, err := s.repo.Store(log.WithContext(ctx), &Hero{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ID:        xid.New().String(),
		Name:      req.Name,
		Owner:     req.Owner,
	})
	if err != nil {
		return nil, fmt.Errorf("cannot store hero '%s': %w", req.Name, err)
	}
	log.Debug().Msg("created hero")

	return &CreateHeroResponse{
		Hero: *h,
	}, nil
}

// NewService returns [Service] implementation.
func NewService(ctx context.Context, repo Repository) *service {
	log := zerolog.Ctx(ctx)
	return &service{
		log:  log,
		repo: repo,
	}
}
