POSTGRESQL_URL=postgres://postgres:postgres@localhost:5555/postgres?sslmode=disable

.PHONY: setup migrate reset

setup:
	$(MAKE) reset
	$(MAKE) migrate

migrate:
	migrate -database "$(POSTGRESQL_URL)" -path db/migrations up

reset:
	migrate -database "$(POSTGRESQL_URL)" -path db/migrations drop -f

test:
	go test ./...
