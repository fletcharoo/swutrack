# Makefile

default: help

help: ## Show this help.
	@egrep '^(.+)\:\ .*##\ (.+)' ${MAKEFILE_LIST} | sed 's/:.*##/#/' | column -t -c 2 -s '#'

run: ## Run all tests.
	PORT=8080 SHUTDOWN_TIMEOUT=30000000000 go run ./services/swutrack