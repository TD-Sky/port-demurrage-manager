package routers

import (
	"github.com/gin-gonic/gin"
	"service/handlers"
	"service/middleware"
)

func General(server *gin.Engine) {
	authm := middleware.Authenticate()

	// 入库信息的路由
	server.POST("/store", authm, handlers.Post_store)
	server.GET("/store", authm, handlers.Get_store)
	server.PUT("/store", authm, handlers.Put_store)

	// 出库信息的路由
	server.GET("/load", authm, handlers.Get_loads)
	server.PUT("/load", authm, handlers.Put_load)

	// 特权路由
	users_grp := server.Group("/users")
	{
		users_grp.POST("/login", handlers.Login)
		users_grp.POST("/info", authm, handlers.Info)
	}
}
