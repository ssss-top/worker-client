PROTOC_INCLUDES=-I$(GOPATH)/src -I./proto/

all: proto
.PHONY: all

proto:
	protoc ${PROTOC_INCLUDES} --micro_out=,paths=source_relative:./proto --go_out=,paths=source_relative:./proto ./proto/worker.proto
.PHONY: proto
