package initializers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/urcop/sber-practice-backend/internal/app/dependencies"
	"github.com/urcop/sber-practice-backend/internal/web"
	"github.com/urcop/sber-practice-backend/internal/web/status"
	"github.com/urcop/sber-practice-backend/internal/web/swagger"
)

func SetupRoutes(app *fiber.App, container *dependencies.Container) {
	ctrl := buildRouters(container)

	for i := range ctrl {
		ctrl[i].DefineRouter(app)
	}
}

func buildRouters(container *dependencies.Container) []web.Controller {
	return []web.Controller{
		status.NewStatusController(),
		swagger.NewSwaggerController(),
	}
}
