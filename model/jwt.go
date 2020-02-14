package model

import "github.com/dgrijalva/jwt-go"

type JWTClaims struct {
	jwt.StandardClaims
	Password    string   `json:"password"`
	Username    string   `json:"username"`
	FullName    string   `json:"full_name"`
	Permissions []string `json:"permissions"`
	Role        string   `json:"role"`
}
