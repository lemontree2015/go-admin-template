package sys

import "github.com/dgrijalva/jwt-go"

type UserJwt struct {
	jwt.StandardClaims
	Id       int
	UserName string
	Roles    []string
}
