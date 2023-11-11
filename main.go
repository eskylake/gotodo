package main

import (
	"log"

	"github.com/eskylake/go-todo/controllers"
	"github.com/eskylake/go-todo/initializer"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func dd(arg ...any) {
	log.Fatalln(arg...)
	panic("Die")
}

func init() {
	config, err := initializer.LoadConfig(".")
	if err != nil {
		dd("Failed to load environment variables! \n", err.Error())
	}

	initializer.ConnectDB(&config)
}

func setupRoutes(app *fiber.App) {
	app.Route("/todos", func(router fiber.Router) {
		router.Post("/", controllers.CreateTodo)
		router.Get("", controllers.GetTodos)
	})

	app.Route("/todos/:id", func(router fiber.Router) {
		router.Get("", controllers.GetTodoById)
		router.Patch("", controllers.UpdateTodo)
		router.Delete("", controllers.DeleteTodo)
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		err := initializer.Ping()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "fail",
				"message": "Database connection error",
				"data":    err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "All Healthy",
		})
	})
}

func main() {
	app := fiber.New()
	micro := fiber.New()

	app.Mount("/api", micro)
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowCredentials: true,
	}))

	setupRoutes(micro)

	log.Fatal(app.Listen(":3000"))
}
