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
