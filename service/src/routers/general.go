package routers

import (
	"github.com/gin-gonic/gin"
	"service/handlers"
)

func General(server *gin.Engine) {
	// 入库信息的路由
	server.POST("/store", handlers.Post_store)
	server.GET("/store", handlers.Get_store)
	server.PUT("/store", handlers.Put_store)

	// 特权路由
	privileged(server)
}
