ifneq (,$(wildcard ./.env))
    include .env
    export
endif

.PHONY: setup-dev teardown-dev

setup-dev:
	docker compose -f Backend/docker-compose.yaml up -d

teardown-dev:
	docker compose -f Backend/docker-compose.yaml down -v

start-be-server:
	set BK_SERVER_IDENTITY=webserver && go -C Backend run .

start-be-consumer:
	set BK_SERVER_IDENTITY=consumer && go -C Backend run .
