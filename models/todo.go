package todo

import (
	"github.com/eskylake/go-todo/database"
	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID          uint `gorm:"primarykey" json:"id"`
	Title       uint `json:"title"`
	CompletedAt uint `json:"completed_at"`
}

func GetTodos(c *fiber.Ctx) error {
	db := database.DBConnection

	var todos []Todo

	db.Find(&todos)

	return c.JSON(&todos)
}
