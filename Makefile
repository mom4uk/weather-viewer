.PHONY: tests
include .env
export

tests:
	source .env && gotestsum --format=short-verbose ./tests/...

start:
	go run cmd/main.go

lint:
	golangci-lint run

migrate:
	migrate -path db/migrations \
  		-database $(DATABASE_URL) up

create-migration:
	migrate create -ext sql -dir db/migrations $(name)

