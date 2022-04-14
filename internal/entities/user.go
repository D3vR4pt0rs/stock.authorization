package entities

import "github.com/dgrijalva/jwt-go/v4"

type Profile struct {
	Credentials
	ID      int     `json:"id"`
	Balance float64 `json:"balance"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}
