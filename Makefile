.PHONY: all
all: build

.PHONY: build
build: proto/*/*.proto
	protoc -I=proto --go_out=plugins=grpc,paths=source_relative:proto proto/*/*.proto

.PHONY: doc
doc: proto/*/*.proto
	protoc -I=proto --doc_out=markdown,README.md:docs proto/*/*.proto
