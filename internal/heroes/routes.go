package heroes

import (
	"context"
	"net/http"
	"time"

	"github.com/justinas/alice"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
)

func registerRoutes(
	ctx context.Context,
	mux *http.ServeMux,
	svc Service,
) error {
	c := alice.New()

	c = c.Append(hlog.NewHandler(*zerolog.Ctx(ctx)))
	c = c.Append(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		hlog.FromRequest(r).Info().
			Str("method", r.Method).
			Stringer("url", r.URL).
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).Send()
	}))
	c = c.Append(hlog.RemoteAddrHandler("ip"))
	c = c.Append(hlog.UserAgentHandler("user_agent"))
	c = c.Append(hlog.RefererHandler("referer"))
	c = c.Append(hlog.RequestIDHandler("req_id", "Request-Id"))

	mux.Handle("POST /heroes", c.Then(HandleCreateHero(svc)))

	return nil
}
