package heroes

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	stdlog "log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
)

const serviceName string = "heroes"

// Run runs heroes service with provided configuration.
func Run(cfg *Config) error {
	logLvl, err := zerolog.ParseLevel(cfg.Logger.Level)
	if err != nil {
		return fmt.Errorf("cannot parse logger level: %w", err)
	}
	log := zerolog.New(os.Stdout).With().Timestamp().Str(
		"service",
		serviceName,
	).Logger().Level(logLvl)
	stdlog.SetOutput(log)

	sigCtx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGTERM,
		syscall.SIGINT,
	)
	defer stop()

	if err := configureDB(log.WithContext(context.Background()), cfg); err != nil {
		return fmt.Errorf("cannot configure database: %w", err)
	}
	pool, err := pgxpool.New(
		context.Background(),
		cfg.Postgres.PoolString("search_path="+cfg.Postgres.Schema),
	)
	if err != nil {
		return fmt.Errorf("cannot create pgx pool: %w", err)
	}
	defer pool.Close()
	repo := NewPg(log.WithContext(context.Background()), pool)
	svc := NewService(log.WithContext(context.Background()), repo)

	mux := http.NewServeMux()
	if err := registerRoutes(log.WithContext(context.Background()), mux, svc); err != nil {
		return fmt.Errorf("cannot register mux routes: %w", err)
	}

	httpSrv := http.Server{
		Addr:              cfg.HTTP.Addr,
		Handler:           mux,
		ReadTimeout:       cfg.HTTP.ReadTimeout,
		ReadHeaderTimeout: cfg.HTTP.ReadHeaderTimeout,
		WriteTimeout:      cfg.HTTP.WriteTimeout,
		IdleTimeout:       cfg.HTTP.IdleTimeout,
		MaxHeaderBytes:    cfg.HTTP.MaxHeaderBytes,
	}

	g, gCtx := errgroup.WithContext(sigCtx)

	g.Go(func() error {
		<-gCtx.Done()
		log.Info().Msg("closing postgres connection pool")
		pool.Close()
		log.Info().Msg("closed postgres connection pool")
		return nil
	})

	g.Go(func() error {
		if cfg.HTTP.TLS.CertFile != "" && cfg.HTTP.TLS.KeyFile != "" {
			log.Info().Msgf("listening for https on %s", httpSrv.Addr)
			if err := httpSrv.ListenAndServeTLS(
				cfg.HTTP.TLS.CertFile,
				cfg.HTTP.TLS.KeyFile,
			); err != nil && err != http.ErrServerClosed {
				return fmt.Errorf("cannot listen and serve tls http server: %w", err)
			}
		} else {
			log.Info().Msgf("listening for http on %s", httpSrv.Addr)
			if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				return fmt.Errorf("cannot listen and serve http server: %w", err)
			}
		}
		return nil
	})

	g.Go(func() error {
		<-gCtx.Done()
		log.Info().Msg("shutting down http server")
		shtdCtx, cancel := context.WithTimeout(
			context.Background(),
			cfg.HTTP.GracefulShutdownTimeout,
		)
		defer cancel()
		if err := httpSrv.Shutdown(shtdCtx); err != nil {
			return fmt.Errorf("cannot shutdown http server: %w", err)
		}
		log.Info().Msg("http server shutdown completed")
		return nil
	})

	if err := g.Wait(); err != nil {
		return fmt.Errorf("cannot wait error group: %w", err)
	}

	return nil
}

func configureDB(ctx context.Context, cfg *Config) error {
	conn, err := pgx.Connect(ctx, cfg.Postgres.String())
	if err != nil {
		return fmt.Errorf("cannot connect to postgres: %w", err)
	}
	if err := createSchema(ctx, conn, cfg.Postgres.Schema); err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}
	if err := migrateUp(ctx, cfg.Postgres.URL("search_path="+cfg.Postgres.Schema)); err != nil {
		return fmt.Errorf("cannot migrate up: %w", err)
	}

	return nil
}
