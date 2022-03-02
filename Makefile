PKGS ?= $(shell go list ./...)

test:
	go test -ldflags="-extldflags=-Wl,--allow-multiple-definition" -race -covermode=atomic -coverprofile coverage.out -v ${PKGS}

build:
	go build -o "app" -ldflags="-extldflags=-Wl,--allow-multiple-definition" -tags musl ./

run:
	go run -ldflags="-extldflags=-Wl,--allow-multiple-definition" ./

interface:
	grpcui --plaintext localhost:3000

generate-proto: imports
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    grpc/proto/api.proto

imports:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
