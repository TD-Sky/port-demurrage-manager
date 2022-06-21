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

var sf_node, _ = snowflake.NewNode(114)

func Get_loads(ctx *gin.Context) {
	db, _ := ctx.Get("db")
	var body models.Body

	loads := dba.Select_loads(db.(*sqlx.DB))
	fees, err := predict(db.(*sqlx.DB), loads)

	if err != nil {
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
	var shipping_order models.PostShippingOrder

	ctx.ShouldBind(&shipping_order)

	dba.Insert_shipping_order(db.(*sqlx.DB), shipping_order, sf_node.Generate())

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

func Get_stat_load_map(ctx *gin.Context) {
	db, _ := ctx.Get("db")
	var body models.Body
	var load_map models.StatLoadMap

	loads := dba.Select_loads(db.(*sqlx.DB))
	fees, err := predict(db.(*sqlx.DB), loads)

    // 没数据，loads就是nil了
	if err != nil || loads == nil {
		body = models.Make_Body(19198)
	} else {
		body = models.Make_Body(20000)

		for i := 0; i < len(loads); i++ {
			loads[i].Fee = fees[i]
		}

		load_map = models.To_StatLoadMap(loads)

		body.Set_data("load_map", load_map)
	}

	ctx.JSON(http.StatusOK, body)
}
