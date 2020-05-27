package main

import (
	"context"
	"errors"

	"github.com/gassara-kys/grpc-sample/proto/greeting"
)

type greetingService struct{}

func (g *greetingService) SayHello(ctx context.Context, message *greeting.HelloRequest) (*greeting.HelloResponse, error) {
	if message.Name == "" {
		return nil, errors.New("No messages")
	}
	return &greeting.HelloResponse{
		Message: message.Name + " hello.",
	}, nil
}
