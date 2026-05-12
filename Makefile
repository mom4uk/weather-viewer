.PHONY: tests

dev:
	docker compose \
      -f docker-compose.yml \
      -f docker-compose.dev.yml \
      up -d

prod:
	docker compose \
		-f docker-compose.yml \
		-f docker-compose.prod.yml \
		up -d --build

tests:
	gotestsum ./tests/...

start:
	go run cmd/main.go

lint:
	golangci-lint run

migrate:
	migrate -path db/migrations \
  		-database $(DATABASE_URL) up

create-migration:
	migrate create -ext sql -dir db/migrations $(name)

tests-local:
	set -a && source .env && set +a && \
	gotestsum ./tests/...

lint-fix:
	golangci-lint run --fix