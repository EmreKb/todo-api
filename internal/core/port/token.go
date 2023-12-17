package port

import "github.com/EmreKb/todo-api/internal/core/domain"

type TokenService interface {
	GetAccessToken(user *domain.User) string
	ValidateAccessToken(token string) *domain.User
}
