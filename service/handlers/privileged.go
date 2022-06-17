package handlers

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
	"service/auth"
	"service/dba"
	"service/models"
)

const salt = "l9D2owwtENby"

func Login(ctx *gin.Context) {
	db, _ := ctx.Get("db")
	var user models.User
	var body models.Body

	ctx.BindJSON(&user)

	switch user_pass(db.(*sqlx.DB), user) {
	case 0:
		token_string, _ := auth.Gen(user.Name)
		body = models.Make_Body(0)
		body.Set_data("accessToken", token_string)
	case 1:
		body = models.Make_Body(1)
	case 2:
		body = models.Make_Body(2)
	}

	ctx.JSON(http.StatusOK, body)
}

func Info(ctx *gin.Context) {
	username, _ := ctx.Get("username")

	body := models.Make_Body(20000)

	user_info := models.UserInfo{
		Name:  username.(string),
		Roles: []string{"admin"},
	}

	body.Set_data("user", user_info)

	ctx.JSON(http.StatusOK, body)
}

/*
 *  异常处理码
 *  - 0 通过验证
 *  - 1 用户不存在
 *  - 2 密码错误
 */
func user_pass(db *sqlx.DB, user models.User) uint16 {
	var buf bytes.Buffer

	buf.WriteString(user.Password)
	buf.WriteString(salt)

	hash_pswd := fmt.Sprintf("%x", sha256.Sum256(buf.Bytes()))

	correct_pswd, err := dba.Get_password(db, user.Name)

	// 条件顺序固定
	if err != nil {
		return 1
	} else if hash_pswd == correct_pswd {
		return 0
	} else {
		return 2
	}
}
