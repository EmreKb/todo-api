package domain

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	Completed   bool   `json:"completed" gorm:"default:false"`
}

func NewTodo(title, description string) *Todo {
	return &Todo{
		Title:       title,
		Description: description,
	}
}
