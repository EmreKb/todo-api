package service

import (
	"github.com/EmreKb/todo-api/internal/core/domain"
	"github.com/EmreKb/todo-api/internal/core/port"
)

type TodoService struct {
	todoRepository port.TodoRepository
}

func NewTodoService(todoRepository port.TodoRepository) port.TodoService {
	return &TodoService{todoRepository}
}

func (s *TodoService) GetAllTodos(userId uint) ([]domain.Todo, error) {
	return s.todoRepository.FindAllByUserId(userId)
}

func (s *TodoService) CreateTodo(todo *domain.Todo) (*domain.Todo, error) {
	return s.todoRepository.Save(todo)
}

func (s *TodoService) FindById(id, userId uint) (*domain.Todo, error) {
	return s.todoRepository.FindById(id, userId)
}
