package handler

import (
	"clickledger/internal/service"
	"clickledger/internal/utm"
	"net/http"

	"gorm.io/gorm"
)

type SlugHandler struct {
	linkService  *service.LinkService
	clickService *service.ClickService
}

func CreateSlugHandler(db *gorm.DB) *SlugHandler {
	return &SlugHandler{
		linkService:  service.CreateLinkService(db),
		clickService: service.CreateClickService(db),
	}
}

func (h *SlugHandler) RedirectToSlug() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug := r.PathValue("slug")

		if slug == "" {
			http.Error(w, "Empty slug", 400)
			return
		}

		utmData := utm.DecodeUtm(r)

		link, err := h.linkService.GetLinkBySlug(slug)

		if err != nil {
			return
		}

		http.Redirect(w, r, link.Target, 302)

		_, err = h.clickService.AddNewClick(
			link,
			utmData,
			r.RemoteAddr,
			r.UserAgent(),
			r.Referer(),
		)

		if err != nil {
			panic(err.Error())
		}

	}
}
