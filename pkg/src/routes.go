package source

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/")

	api.Get("/tasks", getTasks)
	api.Post("/tasks", createTask)
	api.Put("/tasks/:id", updateTask)
	api.Delete("/tasks/:id", deleteTask)
	app.Get("/swagger/*", swagger.HandlerDefault)
}
