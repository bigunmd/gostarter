APP_NAME=heroes

test.unit:
	go test -tags unit ./... -v

test.integration.docker.start:
	docker compose -f docker/$(APP_NAME)-integration.yaml up -d --remove-orphans

test.integration.docker.stop:
	docker compose -f docker/$(APP_NAME)-integration.yaml down

test.integration:
	go test -tags integration ./internal/$(APP_NAME) -v
