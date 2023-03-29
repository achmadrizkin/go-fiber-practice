package controller

import (
	"go-fiber-practice/domain"
	"go-fiber-practice/model"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type NovelController struct {
	NovelUseCase domain.NovelUseCase
}

func NewNovelController(NovelUseCase domain.NovelUseCase) *NovelController {
	return &NovelController{NovelUseCase}
}

func (n *NovelController) GetNovelAll(ctx *fiber.Ctx) error {
	org, err := n.NovelUseCase.GetAllNovel()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	res := model.Response{StatusCode: http.StatusOK, Message: "Get All Novel Success", Data: org}
	return ctx.Status(http.StatusOK).JSON(res)
}
