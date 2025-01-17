package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kelseyhightower/envconfig"
)

type gatewayConf struct {
	Port            string `default:"8080"`
	GreetingSvcAddr string `required:"true" split_words:"true"`
}

func main() {
	var conf gatewayConf
	err := envconfig.Process("", &conf)
	if err != nil {
		appLogger.Fatal(err.Error())
	}

	ctx := context.Background()
	svc, err := newGreetingService(ctx, conf)
	if err != nil {
		appLogger.Fatal(err.Error())
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(httpLogger)
	r.Get("/greeting", svc.greetingHandler)

	appLogger.Infof("starting http server at :%s", conf.Port)
	err = http.ListenAndServe(":"+conf.Port, r)
	if err != nil {
		appLogger.Fatal(err.Error())
	}
}
