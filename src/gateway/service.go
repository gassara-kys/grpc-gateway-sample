package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gassara-kys/grpc-sample/proto/greeting"
	"google.golang.org/grpc"
)

type gatewayService struct {
	greetingSvcAddr string
	greetingSvcConn *grpc.ClientConn
}

func newGreetingService(ctx context.Context, conf gatewayConf) (gatewayService, error) {
	svc := gatewayService{}
	svc.greetingSvcAddr = conf.GreetingSvcAddr
	if err := mustConnGRPC(ctx, &svc.greetingSvcConn, svc.greetingSvcAddr); err != nil {
		return svc, err
	}
	return svc, nil
}

func mustConnGRPC(ctx context.Context, conn **grpc.ClientConn, addr string) error {
	var err error
	*conn, err = grpc.DialContext(ctx, addr,
		grpc.WithInsecure(),
		grpc.WithTimeout(time.Second*3),
	)
	if err != nil {
		return err
	}
	return nil
}

func (g *gatewayService) greetingHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		writeResponse(w, http.StatusBadRequest, map[string]interface{}{
			"error": "Required `name` parameter.",
		})
		return
	}
	msg, err := greeting.NewGreetingClient(g.greetingSvcConn).SayHello(
		r.Context(),
		&greeting.HelloRequest{
			Name: name,
		})
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	writeResponse(w, http.StatusOK, map[string]interface{}{
		"result": msg,
	})
	return
}

func writeResponse(w http.ResponseWriter, status int, body map[string]interface{}) {
	if body == nil {
		w.WriteHeader(status)
		return
	}
	buf, err := json.Marshal(body)
	if err != nil {
		appLogger.Errorf("Response body JSON marshal error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.WriteHeader(status)
	w.Write(buf)
}
