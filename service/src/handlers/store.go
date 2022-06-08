package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
	"service/auth"
	"service/dba"
	"service/models"
)

func Post_store(ctx *gin.Context) {
	db, _ := ctx.Get("db")
	var store models.PostStore

	ctx.ShouldBind(&store)

	dba.Insert_store(db.(*sqlx.DB), store)

	body := auth.Make_Body(20000)

	ctx.JSON(http.StatusOK, body)
}

func Get_store(ctx *gin.Context) {
	db, _ := ctx.Get("db")
	var stores []models.GetStore

	dba.Select_stores(db.(*sqlx.DB), &stores)

	body := auth.Make_Body(20000)
	body.Set_data("stores", stores)

	ctx.JSON(http.StatusOK, body)
}

func Put_store(ctx *gin.Context) {
	db, _ := ctx.Get("db")
	var store models.PutStore

	ctx.ShouldBind(&store)

	dba.Update_store(db.(*sqlx.DB), store)

	body := auth.Make_Body(20000)

	ctx.JSON(http.StatusOK, body)
}
