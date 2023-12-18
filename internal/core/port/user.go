package port

import "github.com/EmreKb/todo-api/internal/core/domain"

type UserRepository interface {
	Save(user *domain.User) (*domain.User, error)
	FindByUsername(username string) (*domain.User, error)
}

type UserService interface {
	GetUserByUsername(username string) (*domain.User, error)
}