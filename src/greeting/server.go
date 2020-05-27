package main

import (
	"fmt"
	"net"

	"github.com/gassara-kys/grpc-sample/proto/greeting"
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"
)

type greetingConf struct {
	Port string `default:"8081"`
}

func main() {
	var conf greetingConf
	err := envconfig.Process("", &conf)
	if err != nil {
		appLogger.Fatal(err.Error())
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", conf.Port))
	if err != nil {
		appLogger.Fatal(err)
	}

	server := grpc.NewServer()
	greeting.RegisterGreetingServer(server, &greetingService{})
	appLogger.Infof("starting gRPC server at :%s", conf.Port)
	server.Serve(l)
}
