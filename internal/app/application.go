package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/urcop/sber-practice-backend/internal/app/dependencies"
	"github.com/urcop/sber-practice-backend/internal/app/initializers"
	"github.com/urcop/sber-practice-backend/internal/repository"
)

type Application struct{}

func InitApplication(app *fiber.App) {
	initializers.InitEnv()
	initializers.SetupRoutes(app)

	repository.NewExampleRepository()

	_ = dependencies.Container{}

}
