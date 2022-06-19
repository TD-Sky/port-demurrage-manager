package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
	"service/dba"
	"service/models"
	"strconv"
)

func Get_stores(ctx *gin.Context) {
	db, _ := ctx.Get("db")

	stores := dba.Select_stores(db.(*sqlx.DB))

	body := models.Make_Body(20000)
	body.Set_data("stores", stores)

	ctx.JSON(http.StatusOK, body)
}

func Post_store(ctx *gin.Context) {
	db, _ := ctx.Get("db")
	var store models.PostStore

	ctx.ShouldBind(&store)

	dba.Insert_store(db.(*sqlx.DB), store)

	body := models.Make_Body(20000)

	ctx.JSON(http.StatusOK, body)
}

func Put_store(ctx *gin.Context) {
	db, _ := ctx.Get("db")
	var store models.PutStore

	ctx.ShouldBind(&store)

	dba.Update_store(db.(*sqlx.DB), store)

	body := models.Make_Body(20000)

	ctx.JSON(http.StatusOK, body)
}

func Delete_store(ctx *gin.Context) {
	db, _ := ctx.Get("db")

	id, _ := strconv.Atoi(ctx.Param("id"))

	dba.Delete_store(db.(*sqlx.DB), int32(id))

	body := models.Make_Body(20000)

	ctx.JSON(http.StatusOK, body)
}
