package http

import (
	"strconv"

	"github.com/EmreKb/todo-api/internal/core/domain"
	"github.com/EmreKb/todo-api/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

type TodoHandler struct {
	todoService port.TodoService
	userService port.UserService
}

type (
	CreateTodoRequestBody struct {
		Title       string `json:"title" validate:"required,gte=1"`
		Description string `json:"description" validate:"required,gte=1"`
	}
)

func NewTodoHandler(todoService port.TodoService, userService port.UserService) *TodoHandler {
	return &TodoHandler{todoService, userService}
}

func (h *TodoHandler) GetTodos(c *fiber.Ctx) error {
	payload := c.Locals("payload").(*port.Payload)
	user, err := h.userService.GetUserByUsername(payload.Username)

	if err != nil {
		return err
	}

	res, err := h.todoService.GetAllTodos(user.ID)

	if err != nil {
		return err
	}

	return c.Status(200).JSON(res)
}

func (h *TodoHandler) CreateTodo(c *fiber.Ctx) error {
	var body CreateTodoRequestBody
	c.BodyParser(&body)

	payload := c.Locals("payload").(*port.Payload)
	user, err := h.userService.GetUserByUsername(payload.Username)

	if err != nil {
		return err
	}

	todo := domain.NewTodo(body.Title, body.Description)
	todo.UserID = user.ID

	res, err := h.todoService.CreateTodo(todo)

	return c.Status(200).JSON(res)
}

func (h *TodoHandler) FindByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return err
	}

	payload := c.Locals("payload").(*port.Payload)
	user, err := h.userService.GetUserByUsername(payload.Username)

	if err != nil {
		return err
	}

	res, err := h.todoService.FindById(uint(id), user.ID)

	return c.Status(200).JSON(res)
}
