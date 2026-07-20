package service

import (
	"clickledger/internal/repository"

	"gorm.io/gorm"
)

type AnalyticsService struct {
	repo *repository.LinkAnalyticsRepository
}

func CreateAnalyticsService(db *gorm.DB) *AnalyticsService {
	return &AnalyticsService{
		repo: repository.CreateLinkAnalyticsRepo(db),
	}
}

func (s *AnalyticsService) GetTopReferers(linkID uint, top int) ([]repository.TopReferer, error) {
	data, err := s.repo.GetTopReferers(linkID, top)

	if err != nil {
		return nil, err
	}

	return data, nil
}
