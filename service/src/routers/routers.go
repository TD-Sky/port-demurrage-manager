package routers

import (
	"github.com/gin-gonic/gin"
	"service/handlers"
)

func General_router(server *gin.Engine) {
	// 入库信息的路由
	server.POST("/store", handlers.Post_store)
	server.GET("/store", handlers.Get_store)
	server.PUT("/store", handlers.Put_store)

	// 用户
	users_group := server.Group("/users")
	{
		users_group.POST("/login", handlers.User_login)
		users_group.POST("/info", handlers.User_info)
	}
}
