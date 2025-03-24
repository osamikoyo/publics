package service

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct{
	jwt.RegisteredClaims
	ID uint `json:"id"`
	Username string `json:"username"`
}

const key = "hello"

func generateToken(userId uint, username string) (string, error){
	claims := Claims{
		ID:   userId,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // срок действия - 24 часа
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "publics",
			Subject:   "user_access",
			ID:        fmt.Sprintf("%d", rand.Uint32()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	return token.SignedString([]byte(key))
}