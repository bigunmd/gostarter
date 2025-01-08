# Go Starter

Sample of my code lies here.
<!--toc:start-->
- [Go Starter](#go-starter)
  - [Project structure](#project-structure)
  - [Run](#run)
  - [Develop](#develop)
  - [Test](#test)
  - [Configuration](#configuration)
    - [heroes](#heroes)
<!--toc:end-->

## Project structure

- `cmd` - executable services.
- `internal` - internal implementations of domain models,
repositories and use cases with exportable handlers.

## Run

As Go

```bash
go run ./cmd/heroes -c config.yaml
```

## Develop

Generate OAS 3.1

```bash
# swaggo v2 is required
# go install github.com/swaggo/swag/v2/cmd/swag@latest
go generate
```

## Test

TBD

## Configuration

All services can be configured via `ENV` variables or `json/yaml` configuration
files passed as argument to the executable.

### heroes

|ENV|json/yaml|Default|Description|
|---|---|---|---|
|`LOGGER_LEVEL`|`logger.level`|`info`|Define logger base level|
|`HTTP_ADDR`|`http.addr`|`:8080`|Serve http server on that address|
|`HTTP_GRACEFUL_SHUTDOWN_TIMEOUT`|`http.gracefulShutdownTimeout`|`20s`|Timeout for active connections graceful shutdown period|
|`HTTP_READ_TIMEOUT`|`http.readTimeout`|`20s`|Timeout for incoming connection read state|
|`HTTP_READ_HEADER_TIMEOUT`|`http.readHeaderTimeout`|`10s`|Timeout for incoming connection read header state|
|`HTTP_WRITE_TIMEOUT`|`http.writeTimeout`|`20s`|Timeout for incoming connection write state|
|`HTTP_IDLE_TIMEOUT`|`http.idleTimeout`|`20s`|Timeout for incoming connection idle state|
|`HTTP_MAX_HEADER_BYTES`|`http.maxHeaderBytes`|`0`|Maximum header bytes size for incoming connections|
