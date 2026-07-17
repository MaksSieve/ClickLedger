package repository

import (
	"clickledger/internal/model"
	"context"
	"log"

	"gorm.io/gorm"
)

type ClickRepository struct {
	db *gorm.DB
}

func CreateClickRepo(db *gorm.DB) *ClickRepository {
	return &ClickRepository{
		db: db,
	}
}

func (r *ClickRepository) Migrate() error {
	err := r.db.AutoMigrate(&model.Click{})
	if err != nil {
		return err
	}
	log.Default().Println("migration completed")
	return nil
}

func (r *ClickRepository) CreateClick(click *model.Click) (*model.Click, error) {
	result := r.db.Create(click)
	if result.Error != nil {
		return nil, result.Error
	}
	return click, nil
}

func (r *ClickRepository) GetClickCollection(linkId uint) ([]model.Click, error) {
	clicks, err := gorm.G[model.Click](r.db).Where(&model.Click{LinkID: int(linkId)}).Find(context.Background())
	if err != nil {
		return nil, err
	}
	return clicks, nil
}
