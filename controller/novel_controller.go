package controller

import (
	"go-fiber-practice/domain"
	"go-fiber-practice/model"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type NovelController struct {
	NovelUseCase domain.NovelUseCase
}

func NewNovelController(NovelUseCase domain.NovelUseCase) *NovelController {
	return &NovelController{NovelUseCase}
}

func (n *NovelController) GetNovelAll(ctx *fiber.Ctx) error {
	novel, err := n.NovelUseCase.GetAllNovel()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	res := model.Response{StatusCode: http.StatusOK, Message: "Get All Novel Success", Data: novel}
	return ctx.Status(http.StatusOK).JSON(res)
}

func (n *NovelController) GetNovelById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	novel, err := n.NovelUseCase.GetNovelById(idInt)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	res := model.Response{
		StatusCode: http.StatusOK,
		Message:    "Get Novel By Id Success",
		Data:       novel,
	}
	return ctx.Status(http.StatusOK).JSON(res)
}
