package controllers

import (
	"github.com/gin-gonic/gin"
	"monaPanel/common/request"
	"monaPanel/common/response"
	"monaPanel/global"
	"monaPanel/service"
)

// Register 注册接口
func Register(c *gin.Context) {
	var form request.Register

	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := service.UserService.Register(form); err != nil {
		global.Log.Error(err.Error())
		response.ServiceFail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}
