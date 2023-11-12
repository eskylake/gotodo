package routers

import (
	"github.com/eskylake/go-todo/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupTodoRoutes(app *fiber.App) {
	app.Route("/todos", func(router fiber.Router) {
		router.Post("/", controllers.CreateTodo)
		router.Get("", controllers.GetTodos)
	})

	app.Route("/todos/:id", func(router fiber.Router) {
		router.Get("", controllers.GetTodoById)
		router.Patch("", controllers.UpdateTodo)
		router.Delete("", controllers.DeleteTodo)
	})
}
