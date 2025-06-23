# This makefile is expected to be run within the nix envrionment so it has all necessary tooling.

##############################################
# Docker targets for testing production.
##############################################
.PHONY: up
up:
	docker-compose up --build -d

.PHONY: down
down:
	docker-compose down --volumes --remove-orphans

##############################################
# Proto targets.
##############################################
.PHONY: proto
proto:
	@buf generate

.PHONY: proto-clean
proto-clean:
	@rm -rf server/proto
	@rm -rf myncer-web/src/generated_grpc

##############################################
# Common.
##############################################
.PHONY: build
build: server-build web-build

.PHONY: test
test: server-test

##############################################
# Server targets.
##############################################
.PHONY: server-dev
server-dev:
	@cd server && go run .

.PHONY: server-build
server-build:
	@cd server && go build ./...

.PHONY: tidy
tidy:
	@cd server && go mod tidy

.PHONY: server-test
server-test:
	@cd server/ && go test ./...

##############################################
# Database targets.
##############################################
.PHONY: db-up
db-up:
	docker-compose up db -d

.PHONY: db-down
db-down:
	docker-compose down db --volumes --remove-orphans

.PHONY: psql
psql:
	psql "postgres://devuser:devpass@localhost:5432/myncer"

##############################################
# Web app targets.
##############################################
.PHONY: web-dev
web-dev:
	@cd myncer-web && pnpm dev

.PHONY: web-build
web-build:
	@cd myncer-web && pnpm install && pnpm build

##############################################
# Nix targets.
##############################################
.PHONY: nix
nix:
	nix develop
