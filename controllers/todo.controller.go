package controllers

import (
	"strconv"
	"strings"
	"time"

	"github.com/eskylake/go-todo/database"
	todo "github.com/eskylake/go-todo/models"
	"github.com/eskylake/go-todo/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateTodo(c *fiber.Ctx) error {
	var payload *todo.CreateTodoSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := utils.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	newTodo := todo.Todo{
		Title:   payload.Title,
		Content: payload.Content,
	}

	result := database.DB.Create(&newTodo)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "Title already exist, please use another title"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"todo": newTodo}})
}

func GetTodos(c *fiber.Ctx) error {
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var todos []todo.Todo
	results := database.DB.Limit(intLimit).Offset(offset).Find(&todos)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(todos), "todos": todos})
}

func UpdateTodo(c *fiber.Ctx) error {
	todoId := c.Params("id")

	var payload *todo.UpdateTodoSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var todo todo.Todo
	result := database.DB.First(&todo, "id = ?", todoId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No todo with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	updates := make(map[string]interface{})
	if payload.Title != "" {
		updates["title"] = payload.Title
	}
	if payload.Content != "" {
		updates["content"] = payload.Content
	}

	updates["updated_at"] = time.Now()

	database.DB.Model(&todo).Updates(updates)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"todo": todo}})
}

func GetTodoById(c *fiber.Ctx) error {
	todoId := c.Params("id")

	var todo todo.Todo
	result := database.DB.First(&todo, "id = ?", todoId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No todo with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"todo": todo}})
}

func DeleteTodo(c *fiber.Ctx) error {
	todoId := c.Params("id")

	result := database.DB.Delete(&todo.Todo{}, "id = ?", todoId)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No todo with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
