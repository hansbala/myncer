# This makefile is expected to be run within the nix envrionment so it has all necessary tooling.

# Proto constants.
PROTO_DIR := backend/proto

# Openapi constants.
OPENAPI_JAR := thirdparty/openapi-generator-cli.jar
OPENAPI_CMD := java -jar $(OPENAPI_JAR)
OPENAPI_FILE := openapi/api.yaml
OPENAPI_GO_OUT := backend/api/
OPENAPI_GO_PKG_NAME := api

##############################################
# Proto targets.
##############################################
.PHONY: proto
proto:
	protoc \
	  --proto_path=$(PROTO_DIR) \
	  --go_out=paths=source_relative:$(PROTO_DIR) \
	  $(PROTO_DIR)/*.proto

##############################################
# Openapi targets.
##############################################
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
# Backend targets.
##############################################
.PHONY: backend
backend-run:
	@cd backend && go run .

.PHONY: backend-build
backend-build:
	@cd backend && go build ./...

.PHONY: tidy
tidy:
	@cd backend && go mod tidy

##############################################
# Database targets.
##############################################
.PHONY: db-up
db-up:
	docker-compose up --build -d

.PHONY: db-down
db-down:
	docker-compose down

.PHONY: db-clean
db-clean:
	docker-compose down --volumes --remove-orphans

.PHONY: psql
psql:
	psql "postgres://devuser:devpass@localhost:5432/myncer"

##############################################
# Nix targets.
##############################################
.PHONY: nix
nix:
	nix develop
