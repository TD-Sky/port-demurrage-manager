package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
	"service/dba"
	"service/models"
)

func Post_store(ctx *gin.Context) {
	db, _ := ctx.Get("db")
	var store models.PostStore

	ctx.ShouldBind(&store)

	dba.Insert_store(db.(*sqlx.DB), store)

	ctx.Status(http.StatusOK)
}

func Get_store(ctx *gin.Context) {
	db, _ := ctx.Get("db")
	var stores []models.GetStore

    dba.Get_store(db.(*sqlx.DB), &stores)

    ctx.JSON(http.StatusOK, stores)
}
