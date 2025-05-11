SHELL := /bin/bash

setup:
	@./scripts/build-db-image.sh

env:
	@./scripts/build-environment.sh
	
db/up:
	docker compose build database
	docker compose up database -d

db/down:
	docker compose down database

db/migrations/up:
	@./scripts/goose/migrations-up.sh

db/migrations/down:
	@./scripts/goose/migrations-down.sh

db/migrations/create:
	@./scripts/goose/migrations-create.sh
