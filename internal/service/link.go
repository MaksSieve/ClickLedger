package service

import (
	"clickledger/internal/encoding"
	"clickledger/internal/model"
	"clickledger/internal/repository"

	"gorm.io/gorm"
)

type LinkCreationRequest struct {
	Target   string `json:"" validate:"required"`
	Name     string `json:"" validate:"required"`
	LinkType string `json:"" validate:"required"`
}

type LinkServiceInterface interface {
	GetLink(id uint) (*model.Link, error)
	CreateLink(link *model.Link) error
}

type LinkService struct {
	repo *repository.LinkRepository
}

func (s *LinkService) GetLink(id uint) (*model.Link, error) {
	return s.repo.GetLinkByID(id)
}

func (s *LinkService) GetTargetBySlug(slug string) (string, error) {
	link, err := s.repo.GetLinkBySlug(slug)
	if err != nil {
		return "", err
	}
	return link.Target, nil
}

func (s *LinkService) CreateLink(data *LinkCreationRequest) (*model.Link, error) {
	var encodedLink = encoding.EncodeToBase62([]byte(data.Target))
	link, err := s.repo.CreateLink(&model.Link{
		Name:     data.Name,
		Target:   data.Target,
		LinkType: data.LinkType,
		Slug:     encodedLink,
	})

	if err != nil {
		return nil, err
	}

	return link, nil
}

func CreateLinkService(db *gorm.DB) *LinkService {
	return &LinkService{
		repo: repository.CreateLinkRepo(db),
	}
}
