package service

import (
	"clickledger/internal/repository"
	"errors"

	"gorm.io/gorm"
)

type AnalyticsService struct {
	analyticsRepo *repository.LinkAnalyticsRepository
	linkRepo      *repository.LinkRepository
}

func CreateAnalyticsService(db *gorm.DB) *AnalyticsService {
	return &AnalyticsService{
		analyticsRepo: repository.CreateLinkAnalyticsRepo(db),
		linkRepo:      repository.CreateLinkRepo(db),
	}
}

func (s *AnalyticsService) GetTopReferers(linkID uint, top int) ([]repository.TopReferer, error) {
	_, err := s.linkRepo.GetLinkByID(linkID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrLinkNotFound
	}

	data, err := s.analyticsRepo.GetTopReferers(linkID, top)

	if err != nil {
		return nil, err
	}

	return data, nil
}
