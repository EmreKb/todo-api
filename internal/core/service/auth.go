package service

import (
	"errors"

	"github.com/EmreKb/todo-api/internal/core/domain"
	"github.com/EmreKb/todo-api/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

var (
	ErrIncorrectPassword = errors.New("Incorrect password")
)

type AuthService struct {
	userRepository port.UserRepository
	tokenService   port.TokenService
}

func NewAuthService(userRepository port.UserRepository, tokenService port.TokenService) port.AuthService {
	return &AuthService{userRepository, tokenService}
}

func (s *AuthService) Register(user *domain.User) error {
	_, err := s.userRepository.Save(user)

	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) Login(user *domain.User) (interface{}, error) {
	u, err := s.userRepository.FindByUsername(user.Username)

	if err != nil {
		return nil, err
	}

	if u.Password != user.Password {
		return nil, ErrIncorrectPassword
	}

	token, err := s.tokenService.GetAccessToken(u)

	if err != nil {
		return nil, err
	}

	return fiber.Map{"access_token": token}, nil
}
