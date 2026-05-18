.PHONY: tests

dev: up compose-migrate

up:
	docker compose \
		-f docker-compose.yml \
		-f docker-compose.dev.yml \
		up -d --build

prod:
	docker compose \
		-f docker-compose.yml \
		-f docker-compose.prod.yml \
		up -d --build

tests:
	gotestsum --format=short-verbose ./tests/...

start:
	go run cmd/main.go

lint:
	golangci-lint run

migrate:
	migrate -path db/migrations -database "$$DATABASE_URL" up

compose-migrate:
	docker compose exec application \
		sh -lc '/go/bin/migrate -path db/migrations -database "$$DATABASE_URL" up'

compose-tests:
	docker compose exec application \
		gotestsum --format=short-verbose ./tests/...

lint-fix:
	golangci-lint run --fix