package main

import (
	"yogolibrary-gofiber-rest-api/internal/api"
	"yogolibrary-gofiber-rest-api/internal/config"
	"yogolibrary-gofiber-rest-api/internal/connection"
	"yogolibrary-gofiber-rest-api/internal/repository"
	"yogolibrary-gofiber-rest-api/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {

	cnf := config.Get()
	dbConn := connection.GetDatabase(cnf.Database)
	app := fiber.New()

	customerRepostory := repository.NewCustomer(dbConn)

	customerSerivce := service.NewCustomer(customerRepostory)

	api.NewCustomer(app, customerSerivce)

	app.Get("/check", check)

	app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)

}

func check(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON("OK")
}
