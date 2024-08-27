include .envrc

MIGRATION_PATH=./cmd/migrate/migrations
DB_ADDR=postgres://admin:adminpassword@localhost/social?sslmode=disable


.PHONY: migrate-create
migrate-create:
	@migrate create -seq -ext sql -dir $(MIGRATION_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATION_PATH) -database=$(DB_ADDR) up

.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATION_PATH) -database=$(DB_ADDR) down $(filter-out $@,$(MAKECMDGOALS))
