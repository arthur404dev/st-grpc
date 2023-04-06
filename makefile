# Makefile for st-grpc

PROTOC = protoc
PROTOC_GEN_GO = protoc-gen-go
PROTOC_GEN_GO_GRPC = protoc-gen-go-grpc

PROTO_DIR = api/st_grpc
PROTO_FILE = $(PROTO_DIR)/st_grpc.proto
PROTO_GO_OUT = $(PROTO_DIR)

.PHONY: all
all: gen-proto build

.PHONY: gen-proto
gen-proto: $(PROTO_FILE)
	$(PROTOC) --go_out=$(PROTO_GO_OUT) --go_opt=paths=source_relative --go-grpc_out=$(PROTO_GO_OUT) --go-grpc_opt=paths=source_relative $(PROTO_FILE)

.PHONY: build
build:
	go build -v ./...

.PHONY: test
test:
	go test -v ./...

.PHONY: clean
clean:
	rm -f $(PROTO_DIR)/*.pb.go

.PHONY: tidy
tidy:
	go mod tidy
