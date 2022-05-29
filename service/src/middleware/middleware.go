package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func StateContext(db *sqlx.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("db", db)
	}
}
