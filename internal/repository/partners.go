package repository

import (
	"fmt"
	"github.com/urcop/sber-practice-backend/internal/config"
	"github.com/urcop/sber-practice-backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PartnersRepository interface {
	Get() ([]*models.Partners, error)
	Create(partner *models.Partners) error
	GetByID(id string) (*models.Partners, error)
	Update(partner *models.Partners) error
	Delete(id string) error
}

type PartnersRepositoryImpl struct {
	db *gorm.DB
}

func (p PartnersRepositoryImpl) Get() ([]*models.Partners, error) {
	var partners []models.Partners
	p.db.Find(&partners)
	var partnersResponse []*models.Partners

	for _, partner := range partners {
		var places []models.Places
		p.db.Find(&places, "partner_id = ?", partner.PartnerId)

		partnersResponse = append(partnersResponse, CreatePartnersResponse(partner, places))
	}
	return partnersResponse, nil
}

func (p PartnersRepositoryImpl) Create(partner *models.Partners) error {
	result := p.db.Create(&partner)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p PartnersRepositoryImpl) GetByID(id string) (*models.Partners, error) {
	var partner *models.Partners
	result := p.db.Where("partner_id = ?", id).First(&partner)
	if result.Error != nil {
		return nil, result.Error
	}

	var places []models.Places
	p.db.Find(&places, "partner_id = ?", id)
	partner.Places = places

	return partner, nil
}

func (p PartnersRepositoryImpl) Update(partner *models.Partners) error {
	result := p.db.Save(&partner)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p PartnersRepositoryImpl) Delete(id string) error {
	p.db.Where("partner_id = ?", id).Delete(&models.Places{})
	result := p.db.Delete(&models.Partners{}, "partner_id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func NewPartnersRepository() PartnersRepository {
	cfg := config.GetConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.Name, cfg.DB.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	pgSvc := &PartnersRepositoryImpl{db: db}
	err = db.AutoMigrate(&models.Partners{}, &models.Places{})
	if err != nil {
		panic(err)
	}
	return pgSvc
}

func CreatePartnersResponse(partner models.Partners, places []models.Places) *models.Partners {
	return &models.Partners{
		PartnerId:      partner.PartnerId,
		Title:          partner.Title,
		Conditions:     partner.Conditions,
		AdditionalInfo: partner.AdditionalInfo,
		Places:         places,
	}
}
