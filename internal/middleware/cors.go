package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"strings"
)

func AddCors() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodDelete,
			fiber.MethodPut,
			fiber.MethodOptions,
		}, ","),
		AllowHeaders:     "Accept,Authorization,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token",
		AllowCredentials: true,
		MaxAge:           300,
	})
}
