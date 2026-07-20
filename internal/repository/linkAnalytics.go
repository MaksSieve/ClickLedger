package repository

import (
	"clickledger/internal/model"
	"context"

	"gorm.io/gorm"
)

type LinkAnalyticsRepository struct {
	db *gorm.DB
}

func CreateLinkAnalyticsRepo(db *gorm.DB) *LinkAnalyticsRepository {
	return &LinkAnalyticsRepository{
		db: db,
	}
}

type TopReferer struct {
	Referer string `gorm:"column:referer"`
	Count   int    `gorm:"column:count"`
}

func (r *LinkAnalyticsRepository) GetTopReferers(linkID uint, top int) ([]TopReferer, error) {
	var results []TopReferer

	err := gorm.G[model.Click](r.db).Table("clicks").
		Select("count(referer) as count, referer").
		Where("link_id = ?", 1).
		Group("referer").
		Order("count DESC").
		Limit(2).
		Scan(context.Background(), &results)

	if err != nil {
		return nil, err
	}

	return results, nil
}
