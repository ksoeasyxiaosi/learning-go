package routers

import (
	"github.com/gin-gonic/gin"
	"webProject/application/api/controller"
)

func apiRouter(r *gin.RouterGroup) {

	r.GET("/ping", controller.PingHandle)
	r.POST("/user", controller.UserInsertHandle)

}
