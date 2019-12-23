package main

import (
	"github.com/gin-gonic/gin"
	"webProject/conf"
	"webProject/middleware"
	"webProject/routers"
)

func main() {

	r := gin.Default()
	// 加载配置
	conf.Init()
	// 加载全局中间件
	middleware.Init(r)

	// 加载路由
	routers.Init(r)

	// 开启服务
	r.Run(":1900")
}
