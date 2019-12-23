package middleware

import "github.com/gin-gonic/gin"

func Init(e *gin.Engine) {
	// 配置跨域
	e.Use(initCors())

	e.Use(LoggerToFile())

	e.Use(gin.Recovery())

}
