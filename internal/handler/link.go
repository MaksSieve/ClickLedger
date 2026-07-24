package handler

import (
	"clickledger/internal/repository"
	"clickledger/internal/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type LinkHandler struct {
	service *service.LinkService
}

type LinkAnalyticsResponse struct {
	TopReferers []repository.TopReferer `json:"topReferers"`
}

func CreateLinkHandler(db *gorm.DB) *LinkHandler {
	return &LinkHandler{
		service: service.CreateLinkService(db),
	}
}

func (h *LinkHandler) GetLinkById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
		if err != nil {
			HandleError(w, NewErrRequestParameterInvalid("id", "int", err.Error()))
			return
		} else {
			link, err := h.service.GetLink(uint(id))
			if err != nil {
				HandleError(w, err)
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
			HandleError(w, err)
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
			HandleError(w, NewErrRequestParameterInvalid("Body", "not empty", ""))
		}

		defer r.Body.Close()

		validator := validator.New()

		reqData := &service.LinkCreationRequest{}

		err := json.NewDecoder(r.Body).Decode(reqData)

		if err != nil {
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
