package handler

import (
	"clickledger/internal/service"
	"net/http"

	"gorm.io/gorm"
)

type SlugHandler struct {
	service *service.LinkService
}

func CreateSlugHandler(db *gorm.DB) *SlugHandler {
	return &SlugHandler{
		service: service.CreateLinkService(db),
	}
}

func (h *SlugHandler) RedirectToSlug() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug := r.PathValue("slug")

		if slug == "" {
			http.Error(w, "Empty slug", 400)
			return
		}

		target, err := h.service.GetTargetBySlug(slug)

		if err != nil {
			if err.Error() == "record not found" {
				http.Error(w, "Not found", 404)
				return
			}
			http.Error(w, err.Error(), 500)
			return
		}

		http.Redirect(w, r, target, 302)

	}
}
