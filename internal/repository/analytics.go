package repository

import (
	"clickledger/internal/model"
	"context"
	"time"

	"gorm.io/gorm"
)

type AnalyticsRepository struct {
	db *gorm.DB
}

func CreateAnalyticsRepo(db *gorm.DB) *AnalyticsRepository {
	return &AnalyticsRepository{
		db: db,
	}
}

type TopReferer struct {
	Referer string `gorm:"column:referer"`
	Count   int    `gorm:"column:count"`
}

func (r *AnalyticsRepository) GetTopReferers(linkID uint, top int) ([]TopReferer, error) {
	var results []TopReferer

	err := gorm.G[model.Click](r.db).Table("clicks").
		Select("count(referer) as count, referer").
		Where("link_id = ?", linkID).
		Group("referer").
		Order("count DESC").
		Limit(top).
		Scan(context.Background(), &results)

	if err != nil {
		return nil, err
	}

	return results, nil
}

func (r *AnalyticsRepository) GetTotalClicks(from time.Time, to time.Time) (int64, error) {
	count, err := gorm.G[model.Click](r.db).
		Table("clics").
		Where("created_at BETWEEN ? AND ?", from, to).
		Count(context.Background(), "id")

	if err != nil {
		return 0, err
	}

	return count, nil
}
