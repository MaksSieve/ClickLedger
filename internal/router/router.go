package router

import (
	"clickledger/internal/handler"
	"clickledger/internal/middleware"
	"net/http"

	"gorm.io/gorm"
)

func CreateMainRouter(db *gorm.DB) *http.ServeMux {
	router := http.NewServeMux()

	telemetryHandler := handler.CreateTelemetryhandler()

	router.HandleFunc("GET /health", middleware.DefaultChain(telemetryHandler.Health()))

	router.Handle("/api/", createLinkRouter(db))

	return router
}

func createLinkRouter(db *gorm.DB) *http.ServeMux {
	router := http.NewServeMux()
	linkHandler := handler.CreateLinkHandler(db)
	router.HandleFunc("GET /api/link/{id}", middleware.DefaultChain(linkHandler.GetLinkById()))
	router.HandleFunc("POST /api/link/shorten", middleware.DefaultChain(linkHandler.CreateLink()))
	return router
}
