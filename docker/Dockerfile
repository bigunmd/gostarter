ARG GOIMAGE="golang:1.23.3"

# Step 1: Modules caching
FROM ${GOIMAGE} AS modules
ARG GOPRIVATE_USER="__token__"
ARG GOPRIVATE_PAT=""
ARG GOPRIVATE=""
RUN apt-get update && apt-get install --reinstall ca-certificates
# Configure private repo
RUN echo -e machine "${GOPRIVATE}\nlogin ${GOPRIVATE_USER}\npassword ${GOPRIVATE_PAT}\n" > ~/.netrc
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Step 2: Builder
FROM ${GOIMAGE} AS builder

ARG APP_NAME

COPY --from=modules /go/pkg /go/pkg
COPY . /${APP_NAME}
WORKDIR /${APP_NAME}
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -o /bin/${APP_NAME} -ldflags="-w -s" ./cmd/${APP_NAME}

# Step 3: Final
FROM scratch

ARG APP_NAME

COPY --from=builder /bin/${APP_NAME} /${APP_NAME}
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

