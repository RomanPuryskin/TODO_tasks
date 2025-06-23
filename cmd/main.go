package main

// @title TODO_task API
// @version 1.0
// @description API для управления задачами
// @host localhost:3000

import (
	"context"
	"log"

	_ "github.com/TODO_task/docs"
	"github.com/TODO_task/pkg/database"
	source "github.com/TODO_task/pkg/src"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// загрузка данных из файла .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// подключаемся к БД
	database.ConnectDB()
	db := database.DB
	defer db.Close(context.Background())

	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	source.RegisterRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
