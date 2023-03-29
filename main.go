package main

import (
	"go-fiber-practice/config"
	"go-fiber-practice/controller"
	"go-fiber-practice/database"
	"go-fiber-practice/model"
	"go-fiber-practice/repo"
	"go-fiber-practice/router"
	"go-fiber-practice/usecase"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func main() {
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	//Database
	db := database.ConnectionDB(&loadConfig)
	db.AutoMigrate(&model.Novel{})

	// redis
	rdb := database.ConnectionRedis(&loadConfig)

	startServer(db, rdb)
}

func startServer(db *gorm.DB, rdb *redis.Client) {
	app := fiber.New()

	novelRepo := repo.NewNovelRepo(db, rdb)
	novelUseCase := usecase.NewNovelUsecase(novelRepo)
	novelController := controller.NewNovelController(novelUseCase)

	routes := router.NewRouter(app, novelController)

	err := routes.Listen(":3400")
	if err != nil {
		panic(err)
	}
}
