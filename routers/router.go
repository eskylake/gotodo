package routers

import (
	"github.com/eskylake/go-todo/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setDefaults(app *fiber.App) {
	health(app)
}

func health(app *fiber.App) {
	app.Get("/health", func(c *fiber.Ctx) error {
		err := database.Ping()
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

func SetupRoutes(app *fiber.App) {
	setDefaults(app)
	SetupTodoRoutes(app)

}

func Cors() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins:     "http://localhost:4000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowCredentials: true,
	})
}
