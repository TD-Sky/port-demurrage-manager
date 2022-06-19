package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
	"service/dba"
	"service/models"
)

func Get_warehouses(ctx *gin.Context) {
	db, _ := ctx.Get("db")

	warehouses := dba.Select_warehouses(db.(*sqlx.DB))

	body := models.Make_Body(20000)
	body.Set_data("warehouses", warehouses)

	ctx.JSON(http.StatusOK, body)
}

func Post_warehouse(ctx *gin.Context) {
	db, _ := ctx.Get("db")
	var warehouse models.Warehouse

	ctx.ShouldBind(&warehouse)

	dba.Insert_warehouse(db.(*sqlx.DB), warehouse)

	body := models.Make_Body(20000)

	ctx.JSON(http.StatusOK, body)
}

func Put_warehouse(ctx *gin.Context) {
	db, _ := ctx.Get("db")
	var warehouse models.Warehouse

	ctx.ShouldBind(&warehouse)

	dba.Update_warehouse(db.(*sqlx.DB), warehouse)

	body := models.Make_Body(20000)

	ctx.JSON(http.StatusOK, body)
}

func Delete_warehouse(ctx *gin.Context) {
	db, _ := ctx.Get("db")

	house_name := ctx.Param("house_name")

	dba.Delete_warehouse(db.(*sqlx.DB), house_name)

	body := models.Make_Body(20000)

	ctx.JSON(http.StatusOK, body)
}
