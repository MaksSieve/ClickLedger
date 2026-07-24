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
	router.Handle("/", createSlugRouter(db))

	return router
}

func createSlugRouter(db *gorm.DB) *http.ServeMux {
	router := http.NewServeMux()
	slugHandler := handler.CreateSlugHandler(db)
	router.HandleFunc("/{slug}", middleware.DefaultChain(slugHandler.RedirectToSlug()))
	return router
}

func createLinkRouter(db *gorm.DB) *http.ServeMux {
	router := http.NewServeMux()
	linkHandler := handler.CreateLinkHandler(db)
	clickHandler := handler.CreateClickHandler(db)
	dashboardHandler := handler.CreateDashboardHandler(db)

	router.HandleFunc("GET /api/link", middleware.DefaultChain(linkHandler.GetLinks()))
	router.HandleFunc("GET /api/link/{id}", middleware.DefaultChain(linkHandler.GetLinkById()))
	router.HandleFunc("POST /api/link/shorten", middleware.DefaultChain(linkHandler.CreateLink()))
	router.HandleFunc("GET /api/link/{linkId}/clicks", middleware.DefaultChain(clickHandler.GetLinkClicks()))
	router.HandleFunc("GET /api/dashboard", middleware.DefaultChain(dashboardHandler.GetUserDashboard()))
	return router
}
