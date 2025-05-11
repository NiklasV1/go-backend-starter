SHELL := /bin/bash

DEFAULT_GOAL := build

.PHONY:setup env db/up db/down db/build db/remove db/migrations/up db/migrations/down db/migrations/create

format:
	go fmt ./backend/...

lint: format
	go vet ./backend/...

build: lint
	go build -o example-backend ./backend/main.go

clean:
	go clean

setup: env db/build

env:
	@./scripts/build-environment.sh
	
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
