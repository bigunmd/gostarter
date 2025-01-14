package heroes

import (
	"context"
	"sync"

	"github.com/rs/zerolog"
)

var _ Repository = (*inMem)(nil)

type inMem struct {
	sync.Mutex
	log    *zerolog.Logger
	heroes []*Hero
}

// Store implements [Repository].
func (i *inMem) Store(ctx context.Context, hero *Hero) (*Hero, error) {
	log := zerolog.Ctx(ctx).With().Any("hero", hero).Logger()

	log.Debug().Msg("storing hero")

	i.Lock()
	for _, h := range i.heroes {
		if h.Name == hero.Name {
			i.Unlock()
			return nil, ErrAlreadyExists
		}
	}
	i.heroes = append(i.heroes, hero)
	i.Unlock()

	log.Debug().Msg("stored hero")

	return hero, nil
}

// NewInMem returns in memory [Repository] implementation.
func NewInMem(ctx context.Context) *inMem {
	log := zerolog.Ctx(ctx)
	return &inMem{
		log:    log,
		heroes: []*Hero{},
	}
}
