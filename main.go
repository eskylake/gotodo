package main

import (
	"fmt"

	"github.com/eskylake/go-todo/database"
	todo "github.com/eskylake/go-todo/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func dd(arg ...interface{}) {
	fmt.Printf("%v", arg)
	panic("Die")
}

func initDatabase() {
	var err error
	dsn := "host=go-todo-postgres user=pg-user password=pg-password dbname=pg-db port=5432"
	database.DBConnection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		dd("Failed to open database")
	}

	fmt.Println("Database connection done")
	database.DBConnection.AutoMigrate(&todo.Todo{})
	fmt.Println("Migration done")
}

func setupRoutes(app *fiber.App) {
	app.Get("/todos", todo.GetTodos)
	app.Get("/todos/:id", todo.GetTodoById)
	app.Post("/todos", todo.CreateTodo)

	fmt.Println("Setup routes done")
}

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func main() {
	app := fiber.New()
	initDatabase()
	app.Get("/", helloWorld)
	setupRoutes(app)
	if err := app.Listen(":3000"); err != nil {
		dd(err)
	}
}
