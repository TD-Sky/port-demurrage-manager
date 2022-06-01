package handlers

import (
	"github.com/gin-gonic/gin"
	// "github.com/jmoiron/sqlx"
	"net/http"
	// "service/dba"
	"service/models"
)

func User_login(ctx *gin.Context) {
	var user models.User

	ctx.ShouldBind(user)

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
	})
}

func User_info(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}
