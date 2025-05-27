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

# Openapi constants.
OPENAPI_JAR := thirdparty/openapi-generator-cli.jar
OPENAPI_CMD := java -jar $(OPENAPI_JAR)
OPENAPI_FILE := openapi/api.yaml
OPENAPI_GO_OUT := server/api/
OPENAPI_GO_PKG_NAME := api

##############################################
# Proto targets.
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

##############################################
# Openapi targets.
##############################################
.PHONY: openapi
openapi: openapi-go

.PHONY: openapi-clean
openapi-clean: openapi-go-clean

.PHONY: openapi-go
openapi-go:
	mkdir -p $(OPENAPI_GO_OUT)
	$(OPENAPI_CMD) generate \
	  -i $(OPENAPI_FILE) \
	  -g go \
	  -o $(OPENAPI_GO_OUT) \
	  --package-name=$(OPENAPI_GO_PKG_NAME) \
		--global-property=models,modelDocs=false,supportingFiles=utils.go

.PHONY: openapi-go-clean
openapi-go-clean:
	rm -rf $(OPENAPI_GO_OUT)

##############################################
# Server targets.
##############################################
.PHONY: server-run
server-run:
	@cd server && go run .

.PHONY: server-build
server-build:
	@cd server && go build ./...

.PHONY: tidy
tidy:
	@cd server && go mod tidy

##############################################
# Database targets.
##############################################
.PHONY: psql
psql:
	psql "postgres://devuser:devpass@localhost:5432/myncer"

##############################################
# Nix targets.
##############################################
.PHONY: nix
nix:
	nix develop
