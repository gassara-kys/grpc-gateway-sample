.PHONY: all clean build doc run log stop
all: run

clean:
	rm -f proto/*/*.pb.go
	docker-compose down --volumes --rmi all

build: proto/*/*.proto
	protoc -I=proto --go_out=plugins=grpc,paths=source_relative:proto proto/*/*.proto
	docker-compose build

doc: proto/*/*.proto
	protoc -I=proto --doc_out=markdown,README.md:docs proto/*/*.proto

run: build
	docker-compose up -d

log:
	docker-compose logs -f

stop:
	docker-compose down
