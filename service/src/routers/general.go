package routers

import (
	"github.com/gin-gonic/gin"
	"service/middleware"
)

func General(server *gin.Engine) {
	authm := middleware.Authenticate()

	// 入库信息的路由
	store(server, authm)

    // 出库信息的路由
    load(server)

	// 特权路由
	privileged(server, authm)
}
