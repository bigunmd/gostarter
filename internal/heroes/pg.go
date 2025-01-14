package heroes

import (
	"context"
	"errors"
	"fmt"

	"github.com/bigunmd/gostarter/gen/heroes/db"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

var _ Repository = (*pg)(nil)

type pg struct {
	log  *zerolog.Logger
	pool *pgxpool.Pool
}

// Store implements [Repository].
func (p *pg) Store(ctx context.Context, hero *Hero) (*Hero, error) {
	log := zerolog.Ctx(ctx).With().Any("hero", hero).Logger()

	q := db.New(p.pool)

	log.Debug().Msg("storing hero")
	if err := q.InsertHero(ctx, db.InsertHeroParams{
		ID:    hero.ID,
		Name:  hero.Name,
		Owner: hero.Owner,
		CreatedAt: pgtype.Timestamptz{
			Time:             hero.CreatedAt,
			InfinityModifier: pgtype.Finite,
			Valid:            true,
		},
		UpdatedAt: pgtype.Timestamptz{
			Time:             hero.UpdatedAt,
			InfinityModifier: pgtype.Finite,
			Valid:            true,
		},
	}); err != nil {
		var pgxErr *pgconn.PgError
		if errors.As(err, &pgxErr) {
			switch pgxErr.Code {
			case pgerrcode.UniqueViolation:
				return nil, fmt.Errorf("cannot insert hero: %w: %w", ErrAlreadyExists, err)
			}
		}
		return nil, fmt.Errorf("cannot insert hero: %w", err)
	}
	log.Debug().Msg("stored hero")

	return hero, nil
}

// NewPg returns Postgres database [Repository] implementation.
func NewPg(ctx context.Context, pool *pgxpool.Pool) *pg {
	log := zerolog.Ctx(ctx)
	return &pg{
		log:  log,
		pool: pool,
	}
}
