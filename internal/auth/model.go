package auth

import "github.com/golang-jwt/jwt"

type Claims struct {
	UserID    string `json:"userId"`
	UserEmail string `json:"userEmail,omitempty"`
	jwt.StandardClaims
}
