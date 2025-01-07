package heroes

import "time"

// Hero represents hero model.
type Hero struct {
	// CreatedAt defines creation date.
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt defines last update date.
	UpdatedAt time.Time `json:"updatedAt"`
	// ID defines unique identifier.
	ID string `json:"id"`
	// Name defines unique (for owner) hero name.
	Name string `json:"name"`
	// Owner defines hero owner.
	Owner string `json:"owner"`
}
