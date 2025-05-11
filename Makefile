SHELL := /bin/bash

.PHONY:setup env db/up db/down db/build db/remove db/migrations/up db/migrations/down db/migrations/create

setup: env db/build

env:
	@./scripts/build-environment.sh
	
db/up:
	@./scripts/docker/database-up.sh

db/down:
	@./scripts/docker/database-down.sh

db/build:
	@./scripts/docker/build-db-image.sh

db/remove:
	@./scripts/docker/database-remove.sh

db/migrations/up:
	@./scripts/goose/migrations-up.sh

db/migrations/down:
	@./scripts/goose/migrations-down.sh

db/migrations/create:
	@./scripts/goose/migrations-create.sh
