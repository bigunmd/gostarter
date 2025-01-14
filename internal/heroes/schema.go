package heroes

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog"
)

func createSchema(ctx context.Context, conn *pgx.Conn, schemaName string) error {
	log := zerolog.Ctx(ctx).With().Str("schemaName", schemaName).Logger()
	log.Debug().Msg("creating schema")
	sql := "CREATE SCHEMA IF NOT EXISTS " + schemaName
	if _, err := conn.Exec(log.WithContext(ctx), sql); err != nil {
		return fmt.Errorf("cannot execute create schema query: %w", err)
	}
	log.Debug().Msg("created schema")
	return nil
}

func dropSchema(ctx context.Context, conn *pgx.Conn, schemaName string) error {
	log := zerolog.Ctx(ctx).With().Str("schemaName", schemaName).Logger()
	log.Debug().Msg("dropping schema")
	sql := "DROP SCHEMA IF EXISTS " + schemaName + " CASCADE"
	if _, err := conn.Exec(log.WithContext(ctx), sql); err != nil {
		return fmt.Errorf("cannot execute drop schema query: %w", err)
	}
	log.Debug().Msg("dropped schema")
	return nil
}
