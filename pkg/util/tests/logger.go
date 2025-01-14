package tests

import (
	"testing"

	"github.com/rs/zerolog"
)

// SetupTestLogger returns zerolog logger for test environment.
func SetupTestLogger(t *testing.T) zerolog.Logger {
	return zerolog.New(zerolog.NewTestWriter(t))
}
