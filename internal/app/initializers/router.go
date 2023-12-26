package initializers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/urcop/sber-practice-backend/internal/web"
	"github.com/urcop/sber-practice-backend/internal/web/status"
	"github.com/urcop/sber-practice-backend/internal/web/swagger"
)

func SetupRoutes(app *fiber.App) {
	ctrl := buildRouters()

	for i := range ctrl {
		ctrl[i].DefineRouter(app)
	}
}

func buildRouters() []web.Controller {
	return []web.Controller{
		status.NewStatusController(),
		swagger.NewSwaggerController(),
	}
}
