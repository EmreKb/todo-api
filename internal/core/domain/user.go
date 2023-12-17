package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Todos    []Todo `json:"todos" gorm:"foreignKey:ID"`
}

func NewUser(username, password string) *User {
	return &User{
		Username: username,
		Password: password,
	}
}


