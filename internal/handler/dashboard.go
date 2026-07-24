package handler

import (
	"clickledger/internal/service"
	"encoding/json"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type DashboardHandler struct {
	service *service.DashboardService
}

func CreateDashboardHandler(db *gorm.DB) *DashboardHandler {
	return &DashboardHandler{
		service: service.CreateDashboardService(db),
	}
}

func (h *DashboardHandler) GetUserDashboard() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fromString := "2026-07-01"
		from, err := time.Parse("YYYY-MM-DD", fromString)

		if err != nil {
			HandleError(w, NewErrRequestParameterInvalid("from", "YYYY-MM-DD", fromString))
			return
		}

		toString := "2026-07-01"
		to, err := time.Parse("YYYY-MM-DD", toString)

		if err != nil {
			HandleError(w, NewErrRequestParameterInvalid("to", "YYYY-MM-DD", toString))
			return
		}

		granularity := "day"

		if granularity != "hour" && granularity != "day" && granularity != "week" && granularity != "year" {
			HandleError(w, NewErrRequestParameterInvalid("granularity", "day|hour|week|year", granularity))
		}

		if response, err := h.service.GetUserDashboard(from, to, granularity); err != nil {
			http.Error(w, err.Error(), 500)
			return
		} else {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		}
	}
}
