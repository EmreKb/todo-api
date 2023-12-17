package config

import (
	"os"

	"github.com/joho/godotenv"
)

type ENV struct {
	PORT   string
	DB_URL string
}

func NewENV() *ENV {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	return &ENV{
		PORT:   os.Getenv("PORT"),
		DB_URL: os.Getenv("DB_URL"),
	}
}
