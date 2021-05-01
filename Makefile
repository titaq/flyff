.PHONY: build

build: build-proto

build-proto:
	protoc -I./internal/domain --gogoslick_out=Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,plugins=grpc:./internal/domain --proto_path=. --proto_path=$(GOPATH)/src flyff.proto
	protoc -I./pkg/flyff --gogoslick_out=Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,plugins=grpc:./pkg/flyff --proto_path=. --proto_path=$(GOPATH)/src packets_inbound.proto
