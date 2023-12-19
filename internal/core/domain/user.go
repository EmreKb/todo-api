package domain

type User struct {
	GormModel
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Todos    []Todo `json:"todos,omitempty" gorm:"foreignKey:ID"`
}

func NewUser(username, password string) *User {
	return &User{
		Username: username,
		Password: password,
	}
}
