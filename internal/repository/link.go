package repository

import (
	"clickledger/internal/model"
	"context"
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

func (r *LinkRepository) Migrate() error {
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

func (r *LinkRepository) GetLinks() ([]model.Link, error) {
	links, err := gorm.G[model.Link](r.db).Find(context.Background())
	if err != nil {
		return nil, err
	}
	return links, nil
}

func (r *LinkRepository) GetLinkBySlug(slug string) (*model.Link, error) {
	var link model.Link
	result := r.db.Where("slug = ?", slug).Select("id", "target", "slug").First(&link)
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
