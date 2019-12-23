package controller

import (
	"github.com/gin-gonic/gin"
	"webProject/application/api/validate"
	"webProject/model"
	"webProject/service"
	"webProject/utils"
)

func UserInsertHandle(c *gin.Context) {

	var param validate.Insert
	if err := c.ShouldBind(&param); err != nil {
		r := utils.ErrorReturn(10100, err.Error())
		r(c)
	}

	var userModel model.User

	userModel.UserName = param.UserName
	userModel.Nickname = param.NickName
	userModel.Password = param.Password

	var userServer service.User
	if _, err := userServer.Insert(&userModel); err != nil {
		r := utils.ErrorReturn(10100, err.Error())
		r(c)
	}
	r := utils.SuccessReturn("success")
	r(c)
}
