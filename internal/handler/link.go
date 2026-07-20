package handler

import (
	"clickledger/internal/repository"
	"clickledger/internal/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type LinkHandler struct {
	service          *service.LinkService
	analyticsService *service.AnalyticsService
}

type LinkAnalyticsResponse struct {
	TopReferers []repository.TopReferer `json:"topReferers"`
}

func CreateLinkHandler(db *gorm.DB) *LinkHandler {
	return &LinkHandler{
		service:          service.CreateLinkService(db),
		analyticsService: service.CreateAnalyticsService(db),
	}
}

func (h *LinkHandler) GetLinkById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		} else {
			link, err := h.service.GetLink(uint(id))
			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			} else {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(link)
			}
		}
	}
}

func (h *LinkHandler) GetLinks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		link, err := h.service.GetLinks()
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		} else {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(link)
		}

	}
}

func (h *LinkHandler) CreateLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			http.Error(w, "Empty body", 400)
		}

		defer r.Body.Close()

		validator := validator.New()

		reqData := &service.LinkCreationRequest{}

		err := json.NewDecoder(r.Body).Decode(reqData)

		if err != nil {
			log.Println("error:", err)
			http.Error(w, err.Error(), 400)
			return
		}

		err = validator.Struct(reqData)

		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		if link, err := h.service.CreateLink(reqData); err != nil {
			http.Error(w, err.Error(), 500)
			return
		} else {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(link)
		}
	}
}

func (h *LinkHandler) GetLinkAnalytics() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		var topReferersNumber int64
		topReferersNumberParam := r.URL.Query()["topReferersNumber"]

		if len(topReferersNumberParam) == 0 || len(topReferersNumberParam) > 1 {
			topReferersNumber = 5
		} else {
			topReferersNumber, err = strconv.ParseInt(topReferersNumberParam[0], 10, 32)
			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}
		}

		topReferers, err := h.analyticsService.GetTopReferers(uint(id), int(topReferersNumber))

		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		response := &LinkAnalyticsResponse{
			TopReferers: topReferers,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

	}
}
