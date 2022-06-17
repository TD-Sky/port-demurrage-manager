package handlers

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
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
	var body models.Body

	loads := dba.Select_loads(db.(*sqlx.DB))

	if fees, err := predict(db.(*sqlx.DB), loads); err != nil {
		body = models.Make_Body(19198)

		body.Set_message(fmt.Sprint(err))
	} else {
		body = models.Make_Body(20000)

		for i := 0; i < len(loads); i++ {
			loads[i].Fee = fees[i]
		}
	}

	body.Set_data("loads", loads)

	ctx.JSON(http.StatusOK, body)
}

func Put_load(ctx *gin.Context) {
	db, _ := ctx.Get("db")
	var load models.PutLoad

	ctx.ShouldBind(&load)

	dba.Update_load(db.(*sqlx.DB), load)

	body := models.Make_Body(20000)

	ctx.JSON(http.StatusOK, body)
}

func Post_loads(ctx *gin.Context) {
	db, _ := ctx.Get("db")
	var loads postLoads

	ctx.ShouldBind(&loads)

	dba.Insert_loads(db.(*sqlx.DB), loads.Inner, sf_node.Generate())

	body := models.Make_Body(20000)

	ctx.JSON(http.StatusOK, body)
}

func Delete_load(ctx *gin.Context) {
	db, _ := ctx.Get("db")

	id, _ := strconv.Atoi(ctx.Param("id"))

	dba.Delete_load(db.(*sqlx.DB), int32(id))

	body := models.Make_Body(20000)

	ctx.JSON(http.StatusOK, body)
}
