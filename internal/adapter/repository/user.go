package repository

import (
	"errors"

	"github.com/EmreKb/todo-api/internal/core/domain"
	"github.com/EmreKb/todo-api/internal/core/port"
	"gorm.io/gorm"
)

var ErrUsernameAlreadyExists = errors.New("Username already exists")

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) port.UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Save(user *domain.User) (*domain.User, error) {
	err := r.db.Create(user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, ErrUsernameAlreadyExists
		}
		return nil, err
	}

	return user, nil
}
