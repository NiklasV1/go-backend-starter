SHELL := /bin/bash

setup:
	@./scripts/build-environment.sh
	@./scripts/docker/build-db-image.sh

env:
	@./scripts/build-environment.sh
	
db/up:
	@./scripts/docker/database-up.sh

db/down:
	@./scripts/docker/database-down.sh

db/remove:
	@./scripts/docker/database-remove.sh

db/migrations/up:
	@./scripts/goose/migrations-up.sh

db/migrations/down:
	@./scripts/goose/migrations-down.sh

db/migrations/create:
	@./scripts/goose/migrations-create.sh
