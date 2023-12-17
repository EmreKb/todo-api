package port

import "github.com/EmreKb/todo-api/internal/core/domain"

type AuthService interface {
	Register(user *domain.User) error
}
