package repository

import (
	"fmt"
	"github.com/urcop/sber-practice-backend/internal/config"
	"github.com/urcop/sber-practice-backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PartnersRepository interface {
	GetPartners() ([]*models.Partners, error)
	CreatePartner(partner *models.Partners) error
	GetPartnerByID(id string) (*models.Partners, error)
	UpdatePartner(partner *models.Partners) error
	DeletePartner(id string) error
}

type PartnersRepositoryImpl struct {
	db *gorm.DB
}

func (p PartnersRepositoryImpl) GetPartners() ([]*models.Partners, error) {
	var partners []*models.Partners
	result := p.db.Find(&partners)
	if result.Error != nil {
		return nil, result.Error
	}
	return partners, nil
}

func (p PartnersRepositoryImpl) CreatePartner(partner *models.Partners) error {
	result := p.db.Create(&partner)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p PartnersRepositoryImpl) GetPartnerByID(id string) (*models.Partners, error) {
	var partner *models.Partners
	result := p.db.Where("id = ?", id).First(&partner)
	if result.Error != nil {
		return nil, result.Error
	}
	return partner, nil
}

func (p PartnersRepositoryImpl) UpdatePartner(partner *models.Partners) error {
	result := p.db.Save(&partner)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p PartnersRepositoryImpl) DeletePartner(id string) error {
	result := p.db.Delete("id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func NewExampleRepository() PartnersRepository {
	cfg := config.GetConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.Name, cfg.DB.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	pgSvc := &PartnersRepositoryImpl{db: db}
	err = db.AutoMigrate(&models.Partners{})
	if err != nil {
		panic(err)
	}
	return pgSvc
}
