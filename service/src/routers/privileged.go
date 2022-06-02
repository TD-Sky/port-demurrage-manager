package routers

import (
	"github.com/gin-gonic/gin"
	"service/handlers"
)

func privileged(server *gin.Engine, authm gin.HandlerFunc) {
	users_grp := server.Group("/users")
	{
		users_grp.POST("/login", handlers.Login)
		users_grp.POST("/info", authm, handlers.Info)
	}
}
