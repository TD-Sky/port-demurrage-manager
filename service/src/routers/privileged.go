package routers

import (
	"github.com/gin-gonic/gin"
	"service/handlers"
)

func privileged(server *gin.Engine) {
	users_grp := server.Group("/users")
	{
		users_grp.POST("/login", handlers.Login)
	}
}
