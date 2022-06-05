package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
	"service/auth"
	"service/dba"
	"service/models"
)

func Get_loads(ctx *gin.Context) {
	db, _ := ctx.Get("db")
	var loads []models.GetLoad

	dba.Select_loads(db.(*sqlx.DB), &loads)

	for i := 0; i < len(loads); i++ {
		retain_YMD(&loads[i].Load_date)
	}

	body := auth.Make_Body(20000)
	body.Set_data("loads", loads)

	ctx.JSON(http.StatusOK, body.To_json())
}

func Put_load(ctx *gin.Context) {
	db, _ := ctx.Get("db")
	var load models.PutLoad

	ctx.ShouldBind(&load)

	dba.Update_load(db.(*sqlx.DB), load)

	body := auth.Make_Body(20000)

	ctx.JSON(http.StatusOK, body.To_json())
}
