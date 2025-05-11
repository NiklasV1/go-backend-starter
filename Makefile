SHELL := /bin/bash

.DEFAULT_GOAL := start

.PHONY:setup generate/env db/up db/down db/build db/remove db/migrations/up db/migrations/down db/migrations/create format lint build clean start

format:
	cd ./backend && go fmt ./...

lint: format
	cd ./backend && go vet ./...

build: lint
	cd ./backend && go build -o example-backend ./main.go

start: build
	./backend/example-backend

clean:
	cd ./backend && go clean

setup: generate/env db/build

generate/env:
	@./scripts/build-environment.sh

generate/repositories:
	cd ./backend && sqlc generate
	
db/build:
	cd ./database && docker build -t example-postgres-database .

db/up: db/build
	docker compose up database -d

db/down:
	docker compose down database

db/remove:
	docker compose down database -v

db/migrations/up:
	@./scripts/goose/migrations-up.sh

db/migrations/down:
	@./scripts/goose/migrations-down.sh

db/migrations/create:
	@./scripts/goose/migrations-create.sh
