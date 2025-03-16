# Makefile

default: help

help: ## Show this help.
	@egrep '^(.+)\:\ .*##\ (.+)' ${MAKEFILE_LIST} | sed 's/:.*##/#/' | column -t -c 2 -s '#'

dc-up: ## Spins up docker compose.
	docker compose up --build -d

dc-down: ## Shut down docker compose.
	docker compose down