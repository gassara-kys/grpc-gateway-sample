# grpc-sample

## Flow

```
UserAgent  --(http)-->  GatewayService  --(gRPC)-->  GreetingService
```

## Requirements
- gRPC
```bash
go get -u google.golang.org/grpc
```
- protoc
```bash
brew upgrade protobuf
```
- protoc-gen-go
```bash
go get -u github.com/golang/protobuf/protoc-gen-go
```
- make
```bash
brew install make
```
