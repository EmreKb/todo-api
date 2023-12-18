package http

import (
	"github.com/EmreKb/todo-api/internal/core/domain"
	"github.com/EmreKb/todo-api/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService port.AuthService
}

type (
	RegisterRequestBody struct {
		Username string `json:"username" validate:"required,gte=4"`
		Password string `json:"password" validate:"required,gte=6"`
	}
	LoginRequestBody struct {
		Username string `json:"username" validate:"required,gte=4"`
		Password string `json:"password" validate:"required,gte=6"`
	}
)

func NewAuthHandler(authService port.AuthService) *AuthHandler {
	return &AuthHandler{authService}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var body RegisterRequestBody
	c.BodyParser(&body)

	u := domain.NewUser(body.Username, body.Password)
	err := h.authService.Register(u)

	if err != nil {
		return err
	}

	return c.Status(201).JSON(fiber.Map{"message": "User registered successfully."})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var body LoginRequestBody
	c.BodyParser(&body)

	u := domain.NewUser(body.Username, body.Password)
	res, err := h.authService.Login(u)

	if err != nil {
		return err
	}

	return c.Status(200).JSON(res)
}
