package middleware

import (
	"errors"
	"strings"

	"github.com/EmreKb/todo-api/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrInvalidAuthHeader = errors.New("Auth header is invalid")
)

type Headers struct {
	Authorization string `reqHeader:"authorization"`
}

type JwtMiddleware struct {
	tokenService port.TokenService
}

func NewJwtMiddleware(tokenService port.TokenService) *JwtMiddleware {
	return &JwtMiddleware{tokenService}
}

func (m *JwtMiddleware) Middleware(c *fiber.Ctx) error {
	var h Headers

	if err := c.ReqHeaderParser(&h); err != nil {
		return err
	}

	authHeader := h.Authorization

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return ErrInvalidAuthHeader
	}

	s := strings.SplitN(authHeader, " ", 2)
	token := s[1]

	payload, err := m.tokenService.ValidateAccessToken(token)
	if err != nil {
		return err
	}

	c.Locals("payload", payload)

	return c.Next()
}
