package api

import (
	"context"
	"net/http"
	"time"
	"yogolibrary-gofiber-rest-api/domain"
	"yogolibrary-gofiber-rest-api/dto"
	"yogolibrary-gofiber-rest-api/internal/util"

	"github.com/gofiber/fiber/v2"
)

type customerApi struct {
	customerService domain.CustomerService
}

func NewCustomer(app *fiber.App, cusomerService domain.CustomerService) {
	ca := customerApi{
		customerService: cusomerService,
	}

	app.Get("/customers", ca.Index)
	app.Post("/customers", ca.Create)

}

func (ca customerApi) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	res, err := ca.customerService.Index(c)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CrateResponseError(err.Error()))
	}

	return ctx.JSON(dto.CrateResponseOk(res))
}

func (ca customerApi) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateCustomerRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).
			JSON(dto.CrateResponseErrorData("validation failed", fails))
	}

	err := ca.customerService.Create(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CrateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusCreated).
		JSON(dto.CrateResponseOk(req))

}
