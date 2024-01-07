package services

import (
	"github.com/golang-jwt/jwt"
	"github.com/urcop/sber-practice-backend/internal/models"
	"github.com/urcop/sber-practice-backend/internal/repository"
)

type UserServiceImpl struct {
	repos repository.UserRepository
}

type UserService interface {
	Login(user *models.User) (string, error)
	GetUser(email string) (*models.User, error)
}

func NewUserService(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{repos: userRepository}
}

func (u *UserServiceImpl) Login(user *models.User) (string, error) {
	token, err := generateJWT(user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *UserServiceImpl) GetUser(email string) (*models.User, error) {
	return u.repos.GetUser(email)
}

func generateJWT(userEmail string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 0,
			IssuedAt:  0,
		},
		Email: userEmail,
	})

	tokenString, err := token.SigningString()
	if err != nil {
		return "", err
	}

	return tokenString, err
}
