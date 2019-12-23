package routers

import (
	"github.com/gin-gonic/gin"
)

func Init(e *gin.Engine) {

	r := e.Group("/api")
	apiRouter(r)

}
