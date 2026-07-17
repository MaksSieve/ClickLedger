package service

import (
	"clickledger/internal/model"
	"clickledger/internal/repository"
	"clickledger/internal/utm"

	"github.com/mileusna/useragent"
	"gorm.io/gorm"
)

type ClickService struct {
	repo *repository.ClickRepository
}

type UserAgentData struct {
	DeviceType string
	OS         string
	Browser    string
	IsBot      bool
}

func CreateClickService(db *gorm.DB) *ClickService {
	return &ClickService{
		repo: repository.CreateClickRepo(db),
	}
}

func (s *ClickService) AddNewClick(link *model.Link, utmData *utm.UtmData, userIP string, ua string, referer string) (*model.Click, error) {

	userAgentData := useragent.Parse(ua)

	click := &model.Click{
		LinkID:       int(link.ID),
		UserIP:       userIP,
		UserAgentRaw: ua,
		Target:       link.Target,

		UtmContent: model.JSONB{
			"utm_campaign": utmData.Campaign,
			"utm_content":  utmData.Content,
			"utm_medium":   utmData.Medium,
			"utm_source":   utmData.Source,
			"utm_term":     utmData.Term,
		},

		DeviceType: userAgentData.Device,
		OS:         userAgentData.OS,
		Browser:    userAgentData.Name,
		IsBot:      userAgentData.Bot,

		Referer:     referer,
		QueryParams: model.JSONB{},
	}

	if userAgentData.Desktop {
		click.DeviceType = "desktop"
	}

	return s.repo.CreateClick(click)
}

func (s *ClickService) GetAllClicksByLinkID(linkID uint) ([]model.Click, error) {
	return s.repo.GetClickCollection(linkID)
}
