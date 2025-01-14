// Package postgres provides helpers for Postgres connection
// handling and configuration.
package postgres

import (
	"fmt"
	"time"
)

// PostgresConfig represents Postgres connection
// configuration parameters.
type PostgresConfig struct {
	// Host defines connection host.
	Host string `json:"host" yaml:"host" env:"HOST" env-default:"127.0.0.1"`
	// Port defines connection port.
	Port int32 `json:"port" yaml:"port" env:"PORT" env-default:"5432"`
	// SSlMode defines connection SSL mode.
	SSLMode string `json:"sslMode" yaml:"sslMode" env:"SSL_MODE" env-default:"disable"`
	// DB defines database name for connection.
	DB string `json:"db" yaml:"db" env:"DB" env-default:"postgres"`
	// Schema defines schema name used for connections.
	Schema string `json:"schema" yaml:"schema" env:"SCHEMA" env-default:"public"`
	// User defines user used for connection.
	User string `json:"user" yaml:"user" env:"USER" env-default:"postgres"`
	// Password defines user's password used for connection.
	Password string `json:"password" yaml:"password" env:"PASSWORD" env-default:"postgres"`
	// MaxConns defines maximum active connections in connection pool.
	MaxConns int `json:"maxConns" yaml:"maxConns" env:"MAX_CONNS" env-default:"10"`
	// MinConns defines minimum active connections in connection pool.
	MinConns int `json:"minConns" yaml:"minConns" env:"MIN_CONNS" env-default:"1"`
	// MaxConnLifetime defines maximum acquired connections lifeteme.
	MaxConnLifetime time.Duration `json:"maxConnLifetime" yaml:"maxConnLifetime" env:"MAX_CONN_LIFETIME" env-default:"10m"`
	// MaxConnIdleTime defines maximum acquired conndections idle time.
	MaxConnIdleTime time.Duration `json:"maxConnIdleTime" yaml:"maxConnIdleTime" env:"MAX_CONN_IDLE_TIME" env-default:"1m"`
	// HealthCheckPeriod defines connection health check intervals.
	HealthCheckPeriod time.Duration `json:"healthCheckPeriod" yaml:"healthCheckPeriod" env:"HEALTH_CHECK_PERIOD" env-default:"10s"`
}

// PoolString returns configuration as string of parameters for connection pool creation.
// Optional parameters can be passed in a form of "key=value"
// strings.
func (p PostgresConfig) PoolString(opts ...string) string {
	conf := fmt.Sprintf(
		"host=%s port=%v sslmode=%s user=%s password=%s dbname=%s pool_max_conns=%v pool_min_conns=%v pool_max_conn_lifetime=%s pool_max_conn_idle_time=%s pool_health_check_period=%s",
		p.Host,
		p.Port,
		p.SSLMode,
		p.User,
		p.Password,
		p.DB,
		p.MaxConns,
		p.MinConns,
		p.MaxConnLifetime,
		p.MaxConnIdleTime,
		p.HealthCheckPeriod,
	)
	for _, v := range opts {
		conf += fmt.Sprintf(" %s", v)
	}
	return conf
}

// String returns configuration as string of parameters.
// Optional parameters can be passed in a form of "key=value"
// strings.
func (p PostgresConfig) String(opts ...string) string {
	conf := fmt.Sprintf(
		"host=%s port=%v sslmode=%s user=%s password=%s dbname=%s",
		p.Host,
		p.Port,
		p.SSLMode,
		p.User,
		p.Password,
		p.DB,
	)
	for _, v := range opts {
		conf += fmt.Sprintf(" %s", v)
	}
	return conf
}

// URL returns configuration as URL string.
// Optional parameters can be passed in a form of "key=value"
// strings.
func (p PostgresConfig) URL(args ...string) string {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%v/%s?sslmode=%s",
		p.User,
		p.Password,
		p.Host,
		p.Port,
		p.DB,
		p.SSLMode,
	)
	for _, v := range args {
		url += fmt.Sprintf("&%s", v)
	}
	return url
}
