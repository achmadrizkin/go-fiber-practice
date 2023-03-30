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
	var response model.Response

	if err := ctx.BodyParser(&novel); err != nil {
		response = model.Response{StatusCode: http.StatusBadRequest, Message: "request invalid, unable to parse request body"}
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	if novel.Author == "" || novel.Name == "" || novel.Description == "" {
		response = model.Response{StatusCode: http.StatusBadRequest, Message: "request invalid, author,name,description is required"}
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	// save into database
	err := n.NovelUseCase.CreateNovel(novel)
	if err != nil {
		response = model.Response{StatusCode: http.StatusBadRequest, Message: err.Error()}
		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}

	// delete redis db
	_, errResponse := n.NovelUseCase.GetAllNovel()
	if errResponse != nil {
		response = model.Response{StatusCode: http.StatusBadRequest, Message: errResponse.Error()}
		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}

	// insert all novel to redis
	_, errAll := n.NovelUseCase.GetAllNovel()
	if errAll != nil {
		response = model.Response{StatusCode: http.StatusBadRequest, Message: errAll.Error()}
		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}

	response = model.Response{StatusCode: http.StatusOK, Message: "Insert data novel Success"}
	return ctx.Status(http.StatusOK).JSON(response)
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
