# grpc-sample

## Flow

```
UserAgent  --(http)-->  GatewayService  --(gRPC)-->  GreetingService
```

## Requirements
- go v1.13~
- protoc v3.12.1~
- GNU Make v3.81~
- docker CE v19.03~
- docker-compose v1.25~
- gRPC
```bash
go get -u google.golang.org/grpc
```
- protoc-gen-go
```bash
go get -u github.com/golang/protobuf/protoc-gen-go
```
