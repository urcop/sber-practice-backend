package partners

import (
	"github.com/gofiber/fiber/v2"
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
	partners services.PartnersService
}

func NewController(partners services.PartnersService) *Controller {
	return &Controller{partners: partners}
}

func (ctrl *Controller) DefineRouter(app *fiber.App) {

	app.Use(middleware.AddCors())
	group := app.Group("/api/v1/partners")
	{

		group.Get("/", ctrl.GetPartners)
		group.Post("/", ctrl.CreatePartner)
		group.Get("/:id", ctrl.GetPartnerByID)
		group.Put("/", ctrl.UpdatePartner)
		group.Delete("/:id", ctrl.DeletePartner)
	}
}

func (ctrl *Controller) GetPartners(ctx *fiber.Ctx) error {
	partners, err := ctrl.partners.GetPartners()
	if err != nil {
		return render.SendError(ctx, http.StatusInternalServerError, err, "failed to get partners")
	}

	return ctx.Status(http.StatusOK).JSON(partners)
}

func (ctrl *Controller) CreatePartner(ctx *fiber.Ctx) error {
	var partner models.Partners

	err := ctx.BodyParser(&partner)
	if err != nil {
		return render.SendError(ctx, http.StatusInternalServerError, err, "failed to parse body")
	}

	err = ctrl.partners.CreatePartner(&partner)
	if err != nil {
		return render.SendError(ctx, http.StatusInternalServerError, err, "failed to create partner")
	}
	return ctx.Status(http.StatusOK).JSON(partner)
}

func (ctrl *Controller) GetPartnerByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	result, err := ctrl.partners.GetPartnerByID(id)
	if err != nil {
		return render.SendError(ctx, http.StatusInternalServerError, err, "failed to find partner")
	}
	return ctx.Status(http.StatusOK).JSON(result)
}

func (ctrl *Controller) UpdatePartner(ctx *fiber.Ctx) error {
	var partner models.Partners

	err := ctx.BodyParser(&partner)
	if err != nil {
		return render.SendError(ctx, http.StatusInternalServerError, err, "failed to parse body")
	}
	err = ctrl.partners.UpdatePartner(&partner)
	if err != nil {
		return render.SendError(ctx, http.StatusInternalServerError, err, "failed to update partner")
	}
	return ctx.Status(http.StatusOK).JSON(partner)
}

func (ctrl *Controller) DeletePartner(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	err := ctrl.partners.DeletePartner(id)
	if err != nil {
		return render.SendError(ctx, http.StatusInternalServerError, err, "failed to delete partner")
	}

	return render.JSONAPIPayload(ctx, http.StatusOK, "successfully delete partner")
}
