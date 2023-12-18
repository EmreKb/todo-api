package main

import (
	"github.com/EmreKb/todo-api/internal/adapter/handler/http"
	"github.com/EmreKb/todo-api/internal/adapter/handler/http/middleware"
	"github.com/EmreKb/todo-api/internal/adapter/repository"
	"github.com/EmreKb/todo-api/internal/config"
	"github.com/EmreKb/todo-api/internal/core/service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	router := fiber.New()

	app.Mount("/api", router)

	// Load env
	env := config.NewENV()

	// Database Connection
	db := repository.NewPostgresConnection(env.DB_URL)

	// Repositories
	userR := repository.NewUserRepository(db)

	// Services
	tokenS := service.NewTokenService()
	authS := service.NewAuthService(userR, tokenS)

	// Handlers
	authH := http.NewAuthHandler(authS)

	// Route Maps
	router.Post("/auth/register", middleware.ValidateMiddleware[http.RegisterRequestBody], authH.Register)
	router.Post("/auth/login", middleware.ValidateMiddleware[http.LoginRequestBody], authH.Login)

	app.Listen(":" + env.PORT)
}
