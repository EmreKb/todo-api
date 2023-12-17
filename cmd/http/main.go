package main

import (
	"github.com/EmreKb/todo-api/internal/adapter/handler/http"
	"github.com/EmreKb/todo-api/internal/adapter/handler/http/middleware"
	"github.com/EmreKb/todo-api/internal/adapter/repository"
	"github.com/EmreKb/todo-api/internal/core/service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	router := fiber.New()

	app.Mount("/api", router)

	// Database Connection
	db := repository.NewPostgresConnection("postgres://postgres:postgres@localhost:5432/todo")

	// Repositories
	userR := repository.NewUserRepository(db)

	// Services
	authS := service.NewAuthService(userR)

	// Handlers
	authH := http.NewAuthHandler(authS)

	// Route Maps
	router.Post("/auth/register", middleware.ValidateMiddleware[http.RegisterRequestBody], authH.Register)

	app.Listen(":3000")
}
