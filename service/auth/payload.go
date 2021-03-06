package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// JWT ç Payload éšć
type UserPl struct {
	jwt.StandardClaims
	Name string `json:"username"`
}

func make_UserPl(name string) UserPl {
	return UserPl{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    "service",
		},
		name,
	}
}
