package handler

import (
	"fmt"
	"net/http"
)

type TelemetryHandler struct{}

func CreateTelemetryhandler() *TelemetryHandler {
	return &TelemetryHandler{}
}

func (h *TelemetryHandler) Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	}
}
