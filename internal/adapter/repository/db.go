package repository

import (
	"github.com/EmreKb/todo-api/internal/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresConnection(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&domain.User{}, &domain.Todo{}); err != nil {
		panic(err)
	}

	return db
}
