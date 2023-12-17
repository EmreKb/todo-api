package service

import (
	"github.com/EmreKb/todo-api/internal/core/domain"
	"github.com/EmreKb/todo-api/internal/core/port"
)

type AuthService struct {
	userRepository port.UserRepository
}

func NewAuthService(userRepository port.UserRepository) port.AuthService {
	return &AuthService{userRepository}
}

func (s *AuthService) Register(user *domain.User) error {
	_, err := s.userRepository.Save(user)

	if err != nil {
		return err
	}

	return nil
}
