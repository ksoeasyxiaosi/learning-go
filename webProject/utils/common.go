package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webProject/serializer"
)

func SuccessReturn(message string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, serializer.Response{
			Code: http.StatusOK,
			Msg:  message,
		})
	}
}

func ErrorReturn(code int, message string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, serializer.Response{
			Code: code,
			Msg:  message,
		})
	}
}
