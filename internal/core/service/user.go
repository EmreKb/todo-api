package service

import (
	"github.com/EmreKb/todo-api/internal/core/domain"
	"github.com/EmreKb/todo-api/internal/core/port"
)

type UserService struct{
	userRepository port.UserRepository
}

func NewUserService(userRepository port.UserRepository) port.UserService {
	return &UserService{userRepository}
}

func (s *UserService) GetUserByUsername(username string) (*domain.User, error) {
	return s.userRepository.FindByUsername(username)
}
