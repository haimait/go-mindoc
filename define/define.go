package define

import (
	"github.com/golang-jwt/jwt/v4"
)

type M map[string]interface{}

type UserClaim struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var (
	JwtKey = "go-mindoc"
)
