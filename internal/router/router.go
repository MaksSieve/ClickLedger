package router

import (
	"clickledger/internal/handler"
	"log"
	"net/http"
)

type Route struct {
	Path    string
	Handler http.HandlerFunc
}

var router = []Route{
	{Path: "/health", Handler: handler.HealthHandler},
}

func API(port string) {
	for _, route := range router {
		http.HandleFunc(route.Path, route.Handler)
	}

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
