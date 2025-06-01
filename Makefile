GOOSE_DBSTRING="root:mauFJcuf5dhRMQrjj@tcp(127.0.0.1:3306)/shopdevgo"
GOOSE_MIGRATION_DIR ?= sql/schema
GOOSE_DRIVER ?= mysql
APP_NAME = server
dev:
	go run ./cmd/$(APP_NAME)

build:
	docker compose up -d --build
	docker compose ps
stop:
	docker compose down
run:
	docker compose up -d && go run ./cmd/$(APP_NAME)
up:
	docker compose up -d

upse:
	@GOOSE_DRIVER=mysql GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up

downse:
	@GOOSE_DRIVER=mysql GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) down

resetse:
	@GOOSE_DRIVER=mysql GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) reset
sqlgen:
	sqlc generate

.PHONY: dev upse downse resetse up stop build

.PHONY: air

create_migration:
	@goose -dir=$(GOOSE_MIGRATION_DIR) create $(name) sql

up_by_one:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up-by-one