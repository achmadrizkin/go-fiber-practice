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

func (n *NovelController) CreateNovel(ctx *fiber.Ctx) error {
	var novel model.Novel

	if err := ctx.BodyParser(&novel); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "request invalid, unable to parse request body"})
	}

	if novel.Author == "" || novel.Name == "" || novel.Description == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "request invalid, author,name,description is required"})
	}

	err := n.NovelUseCase.CreateNovel(novel)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Insert data novel success"})
}

func (n *NovelController) GetNovelAll(ctx *fiber.Ctx) error {
	novel, err := n.NovelUseCase.GetAllNovel()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var res model.Response
	if len(novel) > 0 {
		res = model.Response{StatusCode: http.StatusOK, Message: "Get All Novel Success", Data: novel}
	} else {
		res = model.Response{StatusCode: http.StatusOK, Message: "Get All Novel Success (NULL DATA)"}
	}
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

	var res model.Response
	if novel.Name != "" {
		res = model.Response{StatusCode: http.StatusOK, Message: "Get All Novel Success", Data: novel}
	} else {
		res = model.Response{StatusCode: http.StatusOK, Message: "Get All Novel By Id Success (NULL DATA)"}
	}

	return ctx.Status(http.StatusOK).JSON(res)
}
