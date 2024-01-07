package models

import (
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	jwt.StandardClaims
	Email string `json:"email"`
}
