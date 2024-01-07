package repository

import (
	"errors"
	"fmt"
	"github.com/urcop/sber-practice-backend/internal/config"
	"github.com/urcop/sber-practice-backend/internal/helpers"
	"github.com/urcop/sber-practice-backend/internal/models"
	"github.com/urcop/sber-practice-backend/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUser(email string) (*models.User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (u *UserRepositoryImpl) GetUser(email string) (*models.User, error) {
	var user *models.User
	result := u.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (u *UserRepositoryImpl) CreateUser(user *models.User) error {
	result := u.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *UserRepositoryImpl) seeds(cfg *config.Config) error {
	adminEmail := cfg.Admin.Email
	adminPassword := cfg.Admin.Password

	password, err := helpers.HashPassword(adminPassword)
	if err != nil {
		return err
	}
	user := models.User{
		Email:    adminEmail,
		Password: password,
	}

	_, err = u.GetUser(adminEmail)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		u.CreateUser(&user)
	}

	return nil
}

func NewUserRepository() UserRepository {
	cfg := config.GetConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.Name, cfg.DB.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	pgSvc := &UserRepositoryImpl{db: db}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}

	err = pgSvc.seeds(cfg)
	if err != nil {
		logger.Debug(`failed to migrate "create user"`)
	}
	return pgSvc

}
