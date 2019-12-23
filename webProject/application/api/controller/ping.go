package controller

import (
	"github.com/gin-gonic/gin"
	"time"
	"webProject/utils"
)

func PingHandle(c *gin.Context) {
	r := utils.SuccessReturn(utils.Trans("Success.Ping", time.Now()))
	r(c)
}
