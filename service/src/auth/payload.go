package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// JWT 的 Payload 部分
type UserPl struct {
	jwt.StandardClaims
	Name     string `json:"username"`
}

func make_UserPl(name string) UserPl {
    return UserPl {
        jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
            Issuer: "service",
        },
        name,
    }
}
