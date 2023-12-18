package service

import (
	"errors"
	"time"

	"github.com/EmreKb/todo-api/internal/core/domain"
	"github.com/EmreKb/todo-api/internal/core/port"
	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrJwtVerifyToken = errors.New("JWT token cannot verify")
	ErrJwtParseError  = errors.New("Invalid JWT token format")
)

type TokenService struct{}

func NewTokenService() port.TokenService {
	return &TokenService{}
}

func (s *TokenService) GetAccessToken(user *domain.User) (string, error) {
	payload := port.Payload{Username: user.Username, Exp: time.Now().Add(time.Hour * 24).Unix()}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": payload.Username,
		"exp":      payload.Exp,
	})

	ts, err := token.SignedString([]byte("secretkey"))
	if err != nil {
		return "", err
	}

	return ts, nil
}

func (s *TokenService) ValidateAccessToken(t string) (*port.Payload, error) {
	token, err := jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		return []byte("secretkey"), nil
	})

	if err != nil {
		return nil, ErrJwtParseError
	}

	if !token.Valid {
		return nil, ErrJwtVerifyToken
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &port.Payload{Username: claims["username"].(string)}, nil
	} else {
		return nil, ErrJwtParseError
	}
}
