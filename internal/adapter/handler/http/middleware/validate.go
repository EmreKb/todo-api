package middleware

import (
	"github.com/EmreKb/todo-api/internal/utils"
	"github.com/gofiber/fiber/v2"
)

var Validator = utils.NewValidator()

func ValidateMiddleware[T any](c *fiber.Ctx) error {
	body := new(T)

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	if err := Validator.Struct(body); err != nil {
		return err
	}

	return c.Next()
}
