package routers

import (
	"github.com/gin-gonic/gin"
	"service/handlers"
	"service/middleware"
)

func privileged(server *gin.Engine) {
	users_grp := server.Group("/users")
	{
		users_grp.POST("/login", handlers.Login)
		users_grp.POST("/info", middleware.Authenticate(), handlers.Info)
	}
}
