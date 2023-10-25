package middleware

import "github.com/dgrijalva/jwt-go"

const (
	bearer = "Bearer"
)

var (
	JWTSigningMethod = jwt.SigningMethodHS256
)

type JWTClaims struct {
	jwt.StandardClaims
	Domain   string `json:"domain"`
	ID       string `json:"id"`
	OrigIat  int64  `json:"orig_iat"`
	Username string `json:"username"`
}

type JWTRequest struct {
	ID       string
	Domain   string
	Username string
	ExpireAt int64
}
