package port

import (
	"github.com/EmreKb/todo-api/internal/core/domain"
)

type TokenService interface {
	GetAccessToken(user *domain.User) (string, error)
	ValidateAccessToken(token string) (*Payload, error)
}

type Payload struct {
	Username string
	Exp      int64
}
