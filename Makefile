SHELL := /bin/bash

setup:
	@./scripts/build-environment.sh
	@./scripts/docker/build-db-image.sh
	@./scripts/docker/database-up.sh
	@./scripts/goose/migrations-up.sh
	@./scripts/docker/database-down.sh

env:
	@./scripts/build-environment.sh
	
db/up:
	@./scripts/docker/database-up.sh

db/down:
	@./scripts/docker/database-down.sh

db/migrations/up:
	@./scripts/goose/migrations-up.sh

db/migrations/down:
	@./scripts/goose/migrations-down.sh

db/migrations/create:
	@./scripts/goose/migrations-create.sh
