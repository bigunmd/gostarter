package heroes

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	stdlog "log"

	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
)

const service string = "heroes"

// Run runs heroes service with provided configuration.
func Run(cfg *Config) error {
	logLvl, err := zerolog.ParseLevel(cfg.Logger.Level)
	if err != nil {
		return fmt.Errorf("cannot parse logger level: %w", err)
	}
	log := zerolog.New(os.Stdout).With().Timestamp().Str("service", service).Logger().Level(logLvl)
	stdlog.SetOutput(log)

	sigCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	mux := http.NewServeMux()
	if err := registerRoutes(mux); err != nil {
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
		log.Info().Msgf("listening for http on %s", httpSrv.Addr)
		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return fmt.Errorf("cannot listen and serve http server: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		<-gCtx.Done()
		log.Info().Msg("shutting down http server")
		shtdCtx, cancel := context.WithTimeout(context.Background(), cfg.HTTP.GracefulShutdownTimeout)
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
