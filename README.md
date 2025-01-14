# Go Starter

Sample of my code lies here.

<!--toc:start-->

- [Go Starter](#go-starter)
  - [Project structure](#project-structure)
  - [Run](#run)
  - [Develop](#develop)
    - [Generate](#generate)
    - [Migrations](#migrations)
  - [Test](#test)
    - [Unit](#unit)
    - [Integration](#integration)
  - [Configuration](#configuration) - [heroes](#heroes)
  <!--toc:end-->

## Project structure

- `cmd` - executable services
- `internal` - internal implementations of domain models,
  repositories and use cases with exportable handlers
- `gen` - auto generated code
- `docs` - auto generated OAS specification
- `docker` - docker and docker compose related files
- `pkg` - exportable code

## Run

As Go

```bash
go run ./cmd/heroes -c config.yaml
```

## Develop

### Generate

OAS documentation, sqlc database queries

```bash
# swaggo v2 is required
# go install github.com/swaggo/swag/v2/cmd/swag@latest
#
# sqlc is required
# go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
go generate
```

### Migrations

Add new migrations via `migrate` tool

```bash
# go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
migrate create -dir internal/heroes/migrations -ext .sql -seq create_hero_table
```

## Test

### Unit

Unit tests are built with `unit` tag, so to run them use

```bash
go test -tags=unit ./...
```

Run full suite

```bash
make test.unit
```

### Integration

Integration tests are build with `integration` tag, so to run them use

```bash
go test -tags=integraition ./...
```

Run service based suite

```bash
# start docker containers required for integration checks
make test.integration.docker.start APP_NAME=heroes

# run tests
make test.integration APP_NAME=heroes

# tear down docker containers
make test.integration.docker.stop APP_NAME=heroes
```

## Configuration

All services can be configured via `ENV` variables or `json/yaml` configuration
files passed as argument to the executable.

### heroes

| ENV                              | json/yaml                      | Default     | Description                                                 |
| -------------------------------- | ------------------------------ | ----------- | ----------------------------------------------------------- |
| `LOGGER_LEVEL`                   | `logger.level`                 | `info`      | Define logger base level                                    |
| `HTTP_ADDR`                      | `http.addr`                    | `:8080`     | Serve http server on that address                           |
| `HTTP_GRACEFUL_SHUTDOWN_TIMEOUT` | `http.gracefulShutdownTimeout` | `20s`       | Timeout for active connections graceful shutdown period     |
| `HTTP_READ_TIMEOUT`              | `http.readTimeout`             | `20s`       | Timeout for incoming connection read state                  |
| `HTTP_READ_HEADER_TIMEOUT`       | `http.readHeaderTimeout`       | `10s`       | Timeout for incoming connection read header state           |
| `HTTP_WRITE_TIMEOUT`             | `http.writeTimeout`            | `20s`       | Timeout for incoming connection write state                 |
| `HTTP_IDLE_TIMEOUT`              | `http.idleTimeout`             | `20s`       | Timeout for incoming connection idle state                  |
| `HTTP_MAX_HEADER_BYTES`          | `http.maxHeaderBytes`          | `0`         | Maximum header bytes size for incoming connections          |
| `HTTP_TLS_CERT_FILE`             | `http.tls.certFile`            | `""`        | Path to certificate file                                    |
| `HTTP_TLS_KEY_FILE`              | `http.tls.keyFile`             | `""`        | Path to key file                                            |
| `POSTGRES_HOST`                  | `postgres.host`                | `127.0.0.1` | Postgres connection host                                    |
| `POSTGRES_PORT`                  | `postgres.port`                | `5432`      | Postgres connection port                                    |
| `POSTGRES_SSL_MODE`              | `postgres.sslMode`             | `disable`   | Postgres connection SSL options                             |
| `POSTGRES_DB`                    | `postgres.db`                  | `postgres`  | Postgres connection database                                |
| `POSTGRES_SCHEMA`                | `postgres.schema`              | `public`    | Postgres connection search path schema and migration target |
| `POSTGRES_USER`                  | `postgres.user`                | `postgres`  | Postgres connection user                                    |
| `POSTGRES_PASSWORD`              | `postgres.password`            | `postgres`  | Postgres connection password                                |
| `POSTGRES_MAX_CONNS`             | `postgres.maxConns`            | `10`        | Postgres connection pool max size                           |
| `POSTGRES_MIN_CONNS`             | `postgres.minConns`            | `1`         | Postgres connection pool min size                           |
| `POSTGRES_MAX_CONN_LIFETIME`     | `postgres.maxConnLifetime`     | `10m`       | Postgres pool connection max lifetime                       |
| `POSTGRES_MAX_CONN_IDLE_TIME`    | `postgres.maxConnIdleTime`     | `1m`        | Postgres pool connection max idle time                      |
| `POSTGRES_HEALTH_CHECK_PERIOD`   | `postgres.healthCheckPeriod`   | `10s`       | Postgres pool connection health check period                |
