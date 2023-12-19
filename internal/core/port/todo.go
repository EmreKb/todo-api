package port

import "github.com/EmreKb/todo-api/internal/core/domain"

type TodoService interface {
	GetAllTodos(userId uint) ([]domain.Todo, error)
	CreateTodo(*domain.Todo) (*domain.Todo, error)
	FindById(id, userId uint) (*domain.Todo, error)
}

type TodoRepository interface {
	FindAllByUserId(uint) ([]domain.Todo, error)
	Save(*domain.Todo) (*domain.Todo, error)
	FindById(id, userId uint) (*domain.Todo, error)
}