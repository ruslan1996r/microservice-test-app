package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	Secret []byte
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func NewAuthService(secret string) *AuthService {
	jwtSecret := []byte(secret)

	return &AuthService{
		Secret: jwtSecret,
	}
}

func (s *AuthService) Token() (string, error) {
	claims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(s.Secret)
}

func (s *AuthService) CheckToken(tokenString string) (bool, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return s.Secret, nil
		},
	)
	if err != nil {
		return false, err
	}

	if _, ok := token.Claims.(*Claims); ok && token.Valid {
		return true, nil
	} else {
		return false, err
	}
}
