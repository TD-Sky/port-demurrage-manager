package handlers

import (
	"net/http"
	"service/auth"
	"service/dba"
	"service/models"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Get_loads(ctx *gin.Context) {
	db, _ := ctx.Get("db")
	var loads []models.GetLoad

	dba.Select_loads(db.(*sqlx.DB), &loads)

    body :=  auth.Make_Body(20000)
    body.Set_data("loads", loads)

    ctx.JSON(http.StatusOK, body.To_json())
}
