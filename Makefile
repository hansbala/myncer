PROTO_DIR := backend/proto

.PHONY: backend
backend:
	@cd backend && go run .

.PHONY: proto
proto:
	protoc \
	  --proto_path=$(PROTO_DIR) \
	  --go_out=paths=source_relative:$(PROTO_DIR) \
	  $(PROTO_DIR)/*.proto

psql:
	psql "postgres://devuser:devpass@localhost:5432/myncer"
