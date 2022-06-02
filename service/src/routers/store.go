package routers

import (
	"github.com/gin-gonic/gin"
	"service/handlers"
)

func store(server *gin.Engine, authm gin.HandlerFunc) {
	server.POST("/store", authm, handlers.Post_store)
	server.GET("/store", authm, handlers.Get_store)
	server.PUT("/store", authm, handlers.Put_store)
}
