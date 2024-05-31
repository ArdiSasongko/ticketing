package helper

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

type TokenUseCaseInterface interface {
	GeneraredToken(claims CustomClaims) (string, error)
}

type TokenUseCase struct{}

func NewTokenUseCase() *TokenUseCase {
	return &TokenUseCase{}
}

func (t *TokenUseCase) GeneraredToken(claims CustomClaims) (string, error) {
	plainToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := plainToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return token, nil
}
