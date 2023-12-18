package http

import (
	"github.com/EmreKb/todo-api/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService port.UserService
}

func NewUserHandler(userService port.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) GetMe(c *fiber.Ctx) error {
	payload := c.Locals("payload").(*port.Payload)
	user, err := h.userService.GetUserByUsername(payload.Username)

	if err != nil {
		return err
	}

	return c.Status(200).JSON(user)
}
