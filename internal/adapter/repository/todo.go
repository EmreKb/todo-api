package repository

import (
	"github.com/EmreKb/todo-api/internal/core/domain"
	"github.com/EmreKb/todo-api/internal/core/port"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) port.TodoRepository {
	return &TodoRepository{db}
}

func (r *TodoRepository) FindAllByUserId(userId uint) ([]domain.Todo, error) {
	var todos []domain.Todo
	err := r.db.Where(&domain.Todo{UserID: userId}).Find(&todos).Error
	return todos, err
}

func (r *TodoRepository) Save(todo *domain.Todo) (*domain.Todo, error) {
	err := r.db.Create(todo).Error
	return todo, err
}

func (r *TodoRepository) FindById(id, userId uint) (*domain.Todo, error) {
	var todo *domain.Todo

	filter := &domain.Todo{}
	filter.ID = id
	filter.UserID = userId

	err := r.db.Where(&filter).First(&todo).Error

	return todo, err
}
