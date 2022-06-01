package auth

import "github.com/golang-jwt/jwt/v4"

var secret = []byte("秋水共长天一色")

func Gen(username string) (string, error) {
	user_pl := make_UserPl(username)                            // Payload 部分
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, user_pl) // 创建 JWT 对象

	return token.SignedString(secret)
}

func Parse(token_string string) (*UserPl, error) {
	token, err := jwt.ParseWithClaims(token_string, &UserPl{}, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if user_pl, ok := token.Claims.(*UserPl); ok && token.Valid {
		return user_pl, nil
	} else {
		return nil, err
	}
}
