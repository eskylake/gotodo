package todo

import (
	"time"
)

type Todo struct {
	ID          uint      `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Content     string    `gorm:"type:text;not null" json:"content"`
	CompletedAt time.Time `gorm:"type:timestamp" json:"completed_at,omitempty"`
	UpdatedAt   time.Time `gorm:"type:timestamp" json:"updated_at,omitempty"`
}

type CreateTodoSchema struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type UpdateTodoSchema struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}
