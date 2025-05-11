setup:
	./scripts/build-db-image.sh

env:
	./scripts/build-environment.sh
	
db/up:
	docker compose build database
	docker compose up database -d

db/down:
	docker compose down database

db/migrations/up:
db/migrations/down:
db/migrations/create:
	./scripts/goose-migration-create.sh
