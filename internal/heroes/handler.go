package heroes

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/rs/zerolog/hlog"
)

type key string

const contentType key = "Content-Type"
const location key = "Location"

// CreateHeroRequest represents request for [Hero] creation.
type CreateHeroRequest struct {
	// Hero embeded hero fields.
	Hero
}

// CreateHeroResponse represents response for [Hero] creation.
type CreateHeroResponse struct {
	// Hero embeded hero fields.
	Hero
}

// Service provides service around [Hero] model.
type Service interface {
	// Create creates new [Hero].
	Create(ctx context.Context, req *CreateHeroRequest) (*CreateHeroResponse, error)
}

// HandleCreateHero returns handler func that controls [Hero] creation.
func HandleCreateHero(svc Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &CreateHeroRequest{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log := hlog.FromRequest(r).With().Any("req", req).Logger()
		log.Debug().Msg("creating hero")
		resp, err := svc.Create(log.WithContext(r.Context()), req)
		if err != nil {
			if errors.Is(err, ErrAlreadyExists) {
				http.Error(w, err.Error(), http.StatusConflict)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Debug().Msg("created hero")

		w.Header().Set(string(contentType), "application/json")
    w.Header().Set(string(location), r.URL.Path + "/" + resp.Name)
		w.WriteHeader(http.StatusCreated)
	}
}
