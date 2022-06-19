package routers

import (
	"github.com/gin-gonic/gin"
	"service/handlers"
	"service/middleware"
)

func General(server *gin.Engine) {
	authm := middleware.Authenticate()

	// 场地
	server.GET("/warehouse", authm, handlers.Get_warehouses)
	server.POST("/warehouse", authm, handlers.Post_warehouse)
	server.PUT("/warehouse", authm, handlers.Put_warehouse)
	server.DELETE("/warehouse/:house_name", authm, handlers.Delete_warehouse)

	// 入库信息
	server.GET("/store", authm, handlers.Get_stores)
	server.POST("/store", authm, handlers.Post_store)
	server.PUT("/store", authm, handlers.Put_store)
	server.DELETE("/store/:id", authm, handlers.Delete_store)

	// 出库信息
	server.GET("/load", authm, handlers.Get_loads)
	server.PUT("/load", authm, handlers.Put_load)
	server.POST("/load", authm, handlers.Post_loads)
	server.DELETE("/load/:id", authm, handlers.Delete_load)

	// 用户信息
	users_grp := server.Group("/users")
	{
		users_grp.POST("/login", handlers.Login)
		users_grp.POST("/info", authm, handlers.Info)
	}
}
