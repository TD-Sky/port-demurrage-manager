package handlers

import (
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
	"service/auth"
	"service/dba"
	"service/models"
	"strconv"
)

type postLoads struct {
	Inner []models.PostLoad `json:"loads"`
}

var sf_node, _ = snowflake.NewNode(114)

func Get_loads(ctx *gin.Context) {
	db, _ := ctx.Get("db")
	var loads []models.GetLoad

	dba.Select_loads(db.(*sqlx.DB), &loads)

	body := auth.Make_Body(20000)
	body.Set_data("loads", loads)

	ctx.JSON(http.StatusOK, body)
}

func Put_load(ctx *gin.Context) {
	db, _ := ctx.Get("db")
	var load models.PutLoad

	ctx.ShouldBind(&load)

	dba.Update_load(db.(*sqlx.DB), load)

	body := auth.Make_Body(20000)

	ctx.JSON(http.StatusOK, body)
}

func Post_loads(ctx *gin.Context) {
	db, _ := ctx.Get("db")
	var loads postLoads

	ctx.ShouldBind(&loads)

	dba.Insert_loads(db.(*sqlx.DB), loads.Inner, sf_node.Generate())

	body := auth.Make_Body(20000)

	ctx.JSON(http.StatusOK, body)
}

func Delete_load(ctx *gin.Context) {
	db, _ := ctx.Get("db")

	id, _ := strconv.Atoi(ctx.Param("id"))

	dba.Delete_load(db.(*sqlx.DB), int32(id))

	body := auth.Make_Body(20000)

	ctx.JSON(http.StatusOK, body)
}
