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
	@cd ./database
	${GOPATH}/bin/goose up -env ./database/config.env

db/migrations/down:
	@cd ./database
	${GOPATH}/bin/goose down -env ./database/config.env

db/migrations/create:
	echo "Please enter a name for the new migration:"
	@read migration_name
	@cd ./database
	${GOPATH}/bin/goose create ${migration_name} sql -env ./database/config.env
