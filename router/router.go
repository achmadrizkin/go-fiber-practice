package router

import (
	"go-fiber-practice/controller"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(router *fiber.App, novelController *controller.NovelController) *fiber.App {
	// Define your routes here
	router.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, world!")
	})

	router.Get("/novel", novelController.GetNovelAll)

	return router
}
