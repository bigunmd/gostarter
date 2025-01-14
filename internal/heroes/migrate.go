package heroes

import (
	"context"
	"embed"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/zerolog"
)

//go:embed migrations
var migrations embed.FS

func migrateUp(ctx context.Context, url string) error {
	log := zerolog.Ctx(ctx)
	src, err := iofs.New(migrations, "migrations")
	if err != nil {
		return fmt.Errorf("cannot create source driver: %w", err)
	}
	m, err := migrate.NewWithSourceInstance("iofs", src, url)
	if err != nil {
		return fmt.Errorf("cannot create migrator with source instance: %w", err)
	}
	log.Debug().Msg("applying migrations")
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Warn().Err(err).Msg("already on latest migrations")
			return nil
		}
		return fmt.Errorf("cannot migrate up: %w", err)
	}
	log.Debug().Msg("applied migrations")
	return nil
}
