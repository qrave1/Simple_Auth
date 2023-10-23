package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"rchir7/internal/model"
)

type Auth interface {
	GenerateToken(u model.User) (string, error)
}

type TokenHandler struct {
	Secret []byte
}

func NewTokenHandler(secret []byte) *TokenHandler {
	return &TokenHandler{Secret: secret}
}

func (t TokenHandler) GenerateToken(u model.User) (string, error) {
	j := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"Email":    u.Email,
			"Password": u.Password,
			"Role":     u.Role,
		},
	)

	token, err := j.SignedString(t.Secret)

	if err != nil {
		return "", err
	}

	return token, nil
}
