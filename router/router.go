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
	router.Post("/novel", novelController.CreateNovel)
	router.Get("/novel/:id", novelController.GetNovelById)
	router.Put("/novel/:id", novelController.UpdateNovelById)

	return router
}
