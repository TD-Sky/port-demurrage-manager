package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
	"service/auth"
	"service/models"
)

func StateContext(db *sqlx.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("db", db)
	}
}

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body models.Body

		token_string := ctx.Request.Header.Get("X-Access-Token")

		if token_string == "" {
			body = models.Make_Body(1145)
			ctx.JSON(http.StatusOK, body)
			ctx.Abort()
		}

		if user_pl, err := auth.Parse(token_string); err != nil {
			body = models.Make_Body(1145)
			ctx.JSON(http.StatusOK, body)
			ctx.Abort()
		} else {
			ctx.Set("username", user_pl.Name)
		}
	}
}
