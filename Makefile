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

# Proto constants.
PROTO_DIR := server/proto

##############################################
# Proto targets.
# TODO: Migrate these to use `buf`.
##############################################
.PHONY: proto
proto: proto-go

.PHONY: proto-clean
proto-clean: proto-go-clean

.PHONY: proto-go
proto-go:
	protoc \
	  --proto_path=$(PROTO_DIR) \
	  --go_out=paths=source_relative:$(PROTO_DIR) \
	  $(PROTO_DIR)/*.proto

.PHONY: proto-go-clean
proto-go-clean:
	rm -rf $(PROTO_DIR)/*.pb.go
	rm -rf server/proto

##############################################
# Common.
##############################################
.PHONY: build
build: server-build web-build

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

.PHONY: test
test: server-test

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
