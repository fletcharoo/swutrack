# Makefile

default: help

help: ## Show this help.
	@egrep '^(.+)\:\ .*##\ (.+)' ${MAKEFILE_LIST} | sed 's/:.*##/#/' | column -t -c 2 -s '#'

dc.up: ## Spins up docker compose.
	docker compose up --build -d

dc.down: ## Shut down docker compose.
	docker compose down

## Usage: make goose-create-pg name=new
goose.create-pg: ## Create a new postgres migration.
	goose -dir service/data/db/postgres/migrations -s create $(name) sql

.PHONY: install
install:
	go install github.com/pressly/goose/v3/cmd/goose@latest

test: ## Run all tests.
	cd service; go test -count 1 ./...