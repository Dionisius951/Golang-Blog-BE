package helper

import (
	"fmt"
	"time"

	"github.com/Dionisius951/Golang-Blog-BE/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

var mySigningKey = []byte("Secret")

func GenerateToken(user *models.Users) (string, error) {
	claims := models.MyCustomClaims{
		Id:   user.Id,
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	return ss, err
}

func ValidateToken(tokenString string) (any, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.MyCustomClaims{}, func(token *jwt.Token) (any, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	claims, ok := token.Claims.(*models.MyCustomClaims)

	if !ok || !token.Valid {
		fmt.Printf("Unauthorized")
	}

	return claims, nil
}
