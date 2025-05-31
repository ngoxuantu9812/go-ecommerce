GOOSE_DBSTRING="root:mauFJcuf5dhRMQrjj@tcp(127.0.0.1:3306)/shopdevgo"
GOOSE_MIGRATION_DIR ?= sql/schema

APP_NAME = server
dev:
	go run ./cmd/$(APP_NAME)

run:
	docker compose up -d && go run ./cmd/$(APP_NAME)
kill:
	docker compose kill
up:
	docker compose up -d
down:
	docker compose down

upse:
	@GOOSE_DRIVER=mysql GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up

downse:
	@GOOSE_DRIVER=mysql GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) down

resetse:
	@GOOSE_DRIVER=mysql GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) reset

.PHONY: run upse downse resetse

.PHONY: air

create_migration:
	@goose -dir=$(GOOSE_MIGRATION_DIR) create $(name) sql

up_by_one:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up-by-one