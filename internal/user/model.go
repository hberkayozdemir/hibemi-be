package user

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type UserDTO struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type User struct {
	ID              string    `json:"id" bson:"id"`
	FirstName       string    `json:"firstName" bson:"firstName"`
	LastName        string    `json:"lastName" bson:"lastName""`
	Email           string    `json:"email" bson:"email""`
	Phone           string    `json:"phone" bson:"phone""`
	Password        string    `json:"password" bson:"password"`
	IsEmailActivate bool      `json:"isEmailActivate" bson:"isEmailActivate"`
	UserType        string    `json:"userType" bson:"userType"`
	CreatedAt       time.Time `json:"createdAt" bson:"createdAt"`
}

type UserCredentialsDTO struct {
	Email    string `json:"email""`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

type CustomClaims struct {
	UserType string `json:"userType"`
	jwt.StandardClaims
}

type UsersPageableResponse struct {
	Users []User `json:"users"`
	Page  Page   `json:"page"`
}

type Page struct {
	Number        int `json:"number"`
	Size          int `json:"size",omitempty`
	TotalElements int `json:"totalElements",omitempty"`
	TotalPages    int `json:"totalPages",omitempty"`
}

type ActivationCodeDTO struct {
	Code  string `json:"code"`
	Email string `json:"email"`
}
