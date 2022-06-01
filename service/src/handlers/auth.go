package handlers

import (
	"net/http"
	"service/auth"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var user auth.User
	var body auth.Body

	ctx.BindJSON(&user)

	if user.Correct() {
		token_string, _ := auth.Gen(user.Name)
		body = auth.Make_Body(0, "success")
		body.Set_data("token", token_string)
	} else {
		body = auth.Make_Body(1145, "failed")
	}

	ctx.JSON(http.StatusOK, body.To_json())
}
