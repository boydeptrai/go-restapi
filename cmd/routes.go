package main

import (
	"github.com/boydeptrai/go-restapi/handler"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handler.ListFacts)

	app.Post("/fact", handler.CreateFact)
}
