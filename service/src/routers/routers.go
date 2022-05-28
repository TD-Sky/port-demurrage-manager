package routers

import (
	"github.com/gin-gonic/gin"
	"service/state"
)

func General_router(server *gin.Engine, state *state.State) {
	server.GET("/store", func(ctx *gin.Context) {
		//ctx.JSON(200, gin.H{ "message": "Hello world", })
	})
}
