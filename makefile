ifneq (,$(wildcard ./.env))
    include .env
    export
endif

.PHONY: setup-dev teardown-dev

setup-dev:
	docker compose -f Backend/docker-compose.yaml up -d

teardown-dev:
	docker compose -f Backend/docker-compose.yaml down -v
