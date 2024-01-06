package services

import (
	"github.com/google/uuid"
	"github.com/urcop/sber-practice-backend/internal/models"
	"github.com/urcop/sber-practice-backend/internal/repository"
)

type PartnersService interface {
	GetPartners() ([]*models.Partners, error)
	CreatePartner(partner *models.Partners) error
	GetPartnerByID(id string) (*models.Partners, error)
	UpdatePartner(partner *models.Partners) error
	DeletePartner(id string) error
}

type PartnersServiceImpl struct {
	repos repository.PartnersRepository
}

func (p PartnersServiceImpl) GetPartners() ([]*models.Partners, error) {
	return p.repos.Get()
}

func (p PartnersServiceImpl) CreatePartner(partner *models.Partners) error {
	partner.PartnerId = uuid.New().String()
	return p.repos.Create(partner)
}

func (p PartnersServiceImpl) GetPartnerByID(id string) (*models.Partners, error) {
	return p.repos.GetByID(id)
}

func (p PartnersServiceImpl) UpdatePartner(partner *models.Partners) error {
	return p.repos.Update(partner)
}

func (p PartnersServiceImpl) DeletePartner(id string) error {
	return p.repos.Delete(id)
}

func NewPartnersService(repos repository.PartnersRepository) *PartnersServiceImpl {
	return &PartnersServiceImpl{repos: repos}
}
