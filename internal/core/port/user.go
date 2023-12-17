package port

import "github.com/EmreKb/todo-api/internal/core/domain"

type UserRepository interface {
	Save(user *domain.User) (*domain.User, error)
}
