package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
	"service/dba"
	"service/models"
)

func Get_freight_forwarders(ctx *gin.Context) {
	db, _ := ctx.Get("db")

	freight_forwarders := dba.Select_freight_forwarders(db.(*sqlx.DB))

	body := models.Make_Body(20000)
	body.Set_data("companies", freight_forwarders)

	ctx.JSON(http.StatusOK, body)
}

func Post_freight_forwarder(ctx *gin.Context) {
	db, _ := ctx.Get("db")
	var freight_forwarder models.FreightForwarder

	ctx.ShouldBind(&freight_forwarder)

	dba.Insert_freight_forwarder(db.(*sqlx.DB), freight_forwarder)

	body := models.Make_Body(20000)

	ctx.JSON(http.StatusOK, body)
}

func Put_freight_forwarder(ctx *gin.Context) {
	db, _ := ctx.Get("db")
	var freight_forwarder models.FreightForwarder

	ctx.ShouldBind(&freight_forwarder)

	dba.Update_freight_forwarder(db.(*sqlx.DB), freight_forwarder)

	body := models.Make_Body(20000)

	ctx.JSON(http.StatusOK, body)
}

func Delete_freight_forwarder(ctx *gin.Context) {
	db, _ := ctx.Get("db")
	code := ctx.Param("code")
    var body models.Body

    if err := dba.Delete_freight_forwarder(db.(*sqlx.DB), code); err != nil {
        body = models.Make_Body(3)  // 仍有出货订单，不能删除
    } else {
        body = models.Make_Body(0)
    }

	ctx.JSON(http.StatusOK, body)
}
