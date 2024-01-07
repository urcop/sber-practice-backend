package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/urcop/sber-practice-backend/internal/helpers"
	"github.com/urcop/sber-practice-backend/internal/middleware"
	"github.com/urcop/sber-practice-backend/internal/models"
	"github.com/urcop/sber-practice-backend/internal/services"
	"github.com/urcop/sber-practice-backend/internal/web"
	"github.com/urcop/sber-practice-backend/internal/web/render"
	"net/http"
)

var (
	_ web.Controller = (*Controller)(nil)
)

type Controller struct {
	users services.UserService
}

func NewController(users services.UserService) *Controller {
	return &Controller{users: users}
}

func (ctrl *Controller) DefineRouter(app *fiber.App) {
	app.Use(middleware.AddCors())

	auth := app.Group("/api/v1/auth")
	{
		auth.Post("/login", ctrl.Login)
	}
}

// Login
// @Summary Generate JWT
// @Description Generate JWT
// @Tags users
// @Accept json
// @Produce json
// @Param partner body LoginRequest true "Partner object to be created"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /api/v1/auth/login [post]
func (ctrl *Controller) Login(ctx *fiber.Ctx) error {
	var user *models.User

	err := ctx.BodyParser(&user)
	if err != nil {
		return render.SendError(ctx, http.StatusInternalServerError, err, "can`t parse body")
	}

	userModel, err := ctrl.users.GetUser(user.Email)
	if err != nil {
		return render.SendError(ctx, http.StatusBadRequest, err, "incorrect email or password")
	}

	if ok := helpers.CheckPasswordHash(user.Password, userModel.Password); !ok {
		return render.SendMessage(ctx, http.StatusBadRequest, "invalid email or password")
	}

	token, err := ctrl.users.Login(user)
	if err != nil {
		return render.SendError(ctx, http.StatusInternalServerError, err, "failed to generate token")
	}

	return ctx.Status(http.StatusOK).JSON(LoginResponse{
		AccessToken: token,
		StatusCode:  http.StatusOK,
		Message:     "successfully",
	})
}
