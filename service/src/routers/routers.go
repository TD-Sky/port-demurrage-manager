package routers

import (
	"github.com/gin-gonic/gin"
	"service/handlers"
)

func General_router(server *gin.Engine) {
	server.POST("/store", handlers.Post_store)
	server.GET("/store", handlers.Get_store)
}
