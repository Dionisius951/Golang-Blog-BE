package models

import "github.com/golang-jwt/jwt/v5"

type MyCustomClaims struct {
	Id int `json:"username_id"`
	Role   int `json:"role_id"`
	jwt.RegisteredClaims
}
