package heroes

import "time"

// LoggerConfig represents logger configuration parameters.
type LoggerConfig struct {
	// Level defines logger level.
	Level string `json:"level" yaml:"level" env:"LEVEL" env-default:"info"`
}

// TLSConfig represents TLS configuration parameters.
type TLSConfig struct {
	// CertFile defines path to certificate file.
	CertFile string `json:"certFile" yaml:"certFile" env:"CERT_FILE"`
	// KeyFile defines path to key file.
	KeyFile string `json:"keyFile" yaml:"keyFile" env:"KEY_FILE"`
}

// HTTPConfig represents http server configuration parameters.
type HTTPConfig struct {
	// Addr defines 'net/http' http server parameters.
	Addr string `json:"addr" yaml:"addr" env:"ADDR" env-default:":8080"`
	// GracefulShutdownTimeout defines timeout for 'net/http' http server's shutdown.
	GracefulShutdownTimeout time.Duration `json:"gracefulShutdownTimeout" yaml:"gracefulShutdownTimeout" env:"GRACEFUL_SHUTDOWN_TIMEOUT" env-default:"20s"`
	// ReadTimeout defines read timeout for 'net/http' http server.
	ReadTimeout time.Duration `json:"readTimeout" yaml:"readTimeout" env:"READ_TIMEOUT" env-default:"20s"`
	// ReadHeaderTimeout defines read header timeout for 'net/http' http server.
	ReadHeaderTimeout time.Duration `json:"readHeaderTimeout" yaml:"readHeaderTimeout" env:"READ_HEADER_TIMEOUT" env-default:"10s"`
	// WriteTimeout defines write timeout for 'net/http' http server.
	WriteTimeout time.Duration `json:"writeTimeout" yaml:"writeTimeout" env:"WRITE_TIMEOUT" env-default:"20s"`
	// IdleTimeout defines idle timeout for 'net/http' http server.
	IdleTimeout time.Duration `json:"idleTimeout" yaml:"idleTimeout" env:"IDLE_TIMEOUT" env-default:"20s"`
	// MaxHeaderBytes defines max header bytes for 'net/http' http server.
	MaxHeaderBytes int `json:"maxHeaderBytes" yaml:"maxHeaderBytes" env:"MAX_HEADER_BYTES"`
	// TLS defines TLS parameters.
	TLS TLSConfig `json:"tls" yaml:"tls" env-prefix:"TLS_"`
}

// Config represents configuration parameters for heroes
// services.
type Config struct {
	// Logger defines logger parameters.
	Logger LoggerConfig `json:"logger" yaml:"logger" env-prefix:"LOGGER_"`
	// HTTP defines http server parameters.
	HTTP HTTPConfig `json:"http" yaml:"http" env-prefix:"HTTP_"`
}
