package model

import jwt "github.com/dgrijalva/jwt-go"

//Token struct declaration
type Token struct {
	UserID int
	Name   string
	Email  string
	Role   int
	*jwt.StandardClaims
}
