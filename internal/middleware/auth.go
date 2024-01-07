package middleware

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/urcop/sber-practice-backend/internal/models"
	"net/http"
)

func parseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return nil, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		return claims.Email, nil
	}
	return "", errors.New("invalid access token")
}

func AuthHeader() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		headerParts := ctx.GetReqHeaders()["Authorization"]
		if len(headerParts) != 2 {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"message": "invalid authorization",
			})
		}

		if headerParts[0] != "Bearer" {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"message": "invalid authorization",
			})
		}

		_, err := parseToken(headerParts[1])
		if err != nil {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"message": "invalid authorization",
			})
		}

		return nil
	}

}
