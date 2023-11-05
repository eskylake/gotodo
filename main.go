package main

import (
	"github.com/gofiber/fiber/v2"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World! yay")
}

func main() {
	app := fiber.New()

	app.Get("/", helloWorld)

	app.Listen(":3000")
}
