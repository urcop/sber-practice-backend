package swagger

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/urcop/sber-practice-backend/api"
	"github.com/urcop/sber-practice-backend/internal/web"
)

var _ web.Controller = (*Controller)(nil)

type Controller struct{}

func NewSwaggerController() *Controller {
	return &Controller{}
}

func (c *Controller) DefineRouter(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)
}
