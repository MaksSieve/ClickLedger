package repository

import (
	"clickledger/internal/model"
	"log"

	"gorm.io/gorm"
)

type LinkRepository struct {
	db *gorm.DB
}

func CreateLinkRepo(db *gorm.DB) *LinkRepository {
	return &LinkRepository{
		db: db,
	}
}

func (r *LinkRepository) MigrateLink() error {
	err := r.db.AutoMigrate(&model.Link{})
	if err != nil {
		return err
	}
	log.Default().Println("migration completed")
	return nil
}

func (r *LinkRepository) GetLinkByID(id uint) (*model.Link, error) {
	var link model.Link
	result := r.db.First(&link, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (r *LinkRepository) CreateLink(link *model.Link) (*model.Link, error) {
	result := r.db.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}
