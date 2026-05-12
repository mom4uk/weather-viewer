.PHONY: tests
include .env
export

dev:
	docker compose \
      -f docker-compose.yml \
      -f docker-compose.dev.yml \
      up -d

prod:
	docker compose -f docker-compose.yml -f docker-compose.prod.yml up -d

tests:
	gotestsum --format=short-verbose ./tests/...

start:
	go run cmd/main.go

lint:
	golangci-lint run

migrate:
	migrate -path db/migrations \
  		-database $(DATABASE_URL) up

create-migration:
	migrate create -ext sql -dir db/migrations $(name)

