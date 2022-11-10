package controllers

import (
	"github.com/gin-gonic/gin"
	request2 "monaPanel/app/common/request"
	"monaPanel/app/common/response"
	service2 "monaPanel/app/service"
	"monaPanel/global"
	"strconv"
)

// Register 注册接口
func Register(c *gin.Context) {
	var form request2.Register

	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request2.GetErrorMsg(form, err))
		return
	}

	if err, user := service2.UserService.Register(form); err != nil {
		global.Log.Error(err.Error())
		response.ServiceFail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}

// Login 登录接口
func Login(c *gin.Context) {
	var form request2.Login

	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request2.GetErrorMsg(form, err))
	}

	if err, user := service2.UserService.Login(form); err != nil {
		response.ServiceFail(c, err.Error())
	} else {
		tokenData, err := service2.JwtService.CreateToken(user)
		if err != nil {
			response.ServiceFail(c, err.Error())
			return
		}
		response.Success(c, tokenData)
	}
}

// GetUserInfo 获取用户信息接口
func GetUserInfo(c *gin.Context) {
	userId, err := strconv.Atoi(c.Keys["userId"].(string))
	if err != nil {
		response.ClaimsTokenFail(c)
		return
	}

	err, user := service2.UserService.GetUserInfoById(userId)
	if err != nil {
		response.ServiceFail(c, err.Error())
		return
	}
	response.Success(c, user)
}
