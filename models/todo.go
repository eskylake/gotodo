package todo

import (
	"time"

	"github.com/eskylake/go-todo/database"
	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID          uint      `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	CompletedAt time.Time `gorm:"type:timestamp" json:"completed_at"`
}

func GetTodos(c *fiber.Ctx) error {
	db := database.DBConnection

	var todos []Todo

	db.Find(&todos)

	return c.JSON(&todos)
}

func GetTodoById(c *fiber.Ctx) error {
	var todo Todo
	id := c.Params("id")
	db := database.DBConnection

	err := db.Find(&todo, id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Couldn't find Todo",
			"data":    err,
		})
	}

	return c.JSON(&todo)
}

func CreateTodo(c *fiber.Ctx) error {
	db := database.DBConnection
	todo := new(Todo)

	parseErr := c.BodyParser(todo)
	if parseErr != nil {
		return c.Status(422).JSON(fiber.Map{
			"status":  "error",
			"message": "Data parsing error. Check your data",
			"data":    parseErr,
		})
	}

	todoErr := db.Create(&todo).Error
	if todoErr != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Couldn't create Todo",
			"data":    todoErr,
		})
	}

	return c.JSON(&todo)
}
