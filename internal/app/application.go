package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/urcop/sber-practice-backend/internal/app/dependencies"
	"github.com/urcop/sber-practice-backend/internal/app/initializers"
	"github.com/urcop/sber-practice-backend/internal/repository"
	"github.com/urcop/sber-practice-backend/internal/services"
)

type Application struct{}

func InitApplication(app *fiber.App) {
	initializers.InitEnv()

	partnerRepos := repository.NewPartnersRepository()
	partnerService := services.NewPartnersService(partnerRepos)

	userRepository := repository.NewUserRepository()
	userService := services.NewUserService(userRepository)

	container := &dependencies.Container{
		Partners: partnerService,
		Users:    userService,
	}

	initializers.SetupRoutes(app, container)
}
