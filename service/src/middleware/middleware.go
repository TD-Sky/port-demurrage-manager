package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
	"service/auth"
)

func StateContext(db *sqlx.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("db", db)
	}
}

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body auth.Body

		token_string := ctx.Request.Header.Get("X-Access-Token")

		if token_string == "" {
			body = auth.Make_Body(1145)
			ctx.JSON(http.StatusOK, body.To_json())
			ctx.Abort()
		}

		if user_pl, err := auth.Parse(token_string); err != nil {
			body = auth.Make_Body(1145)
			ctx.JSON(http.StatusOK, body.To_json())
			ctx.Abort()
		} else {
			ctx.Set("username", user_pl.Name)
		}
	}
}
