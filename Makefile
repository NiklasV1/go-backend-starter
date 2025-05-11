setup:
	./scripts/build-db-image.sh

env:
	./scripts/build-environment.sh
	
db/up:
db/down:
db/migrations/up:
db/migrations/down:
db/migrations/create:
	${GOPATH}/bin/goose
