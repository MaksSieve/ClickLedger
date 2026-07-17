package handler

import (
	"clickledger/internal/service"
	"encoding/json"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type ClickHandler struct {
	service *service.ClickService
}

func CreateClickHandler(db *gorm.DB) *ClickHandler {
	return &ClickHandler{
		service: service.CreateClickService(db),
	}
}

func (h *ClickHandler) GetLinkClicks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		linkId, err := strconv.ParseUint(r.PathValue("linkId"), 10, 32)

		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		clicks, err := h.service.GetAllClicksByLinkID(uint(linkId))

		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		} else {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(clicks)
		}
	}
}
