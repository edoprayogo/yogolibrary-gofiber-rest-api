package api

import (
	"context"
	"net/http"
	"time"
	"yogolibrary-gofiber-rest-api/domain"
	"yogolibrary-gofiber-rest-api/dto"

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
