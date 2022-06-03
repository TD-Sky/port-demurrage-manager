package routers

import (
	"service/handlers"

	"github.com/gin-gonic/gin"
)

func load(server *gin.Engine) {
    server.GET("/load", handlers.Get_loads)
}
