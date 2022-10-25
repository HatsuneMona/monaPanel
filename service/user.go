package service

import (
	"errors"
	"monaPanel/common/request"
	"monaPanel/global"
	"monaPanel/models"
	"monaPanel/utils"
)

type userService struct{}

var UserService = new(userService)

// Register 注册
func (userService *userService) Register(params request.Register) (err error, user models.User) {
	mobileCheck := global.DB.Where("mobile = ?", params.Mobile).Select("id").First(&models.User{})

	if mobileCheck.RowsAffected != 0 {
		err = errors.New("手机号已存在")
		return
	}

	loginNameCheck := global.DB.Where("loginName = ?", params.LoginName).Select("id").First(&models.User{})
	if loginNameCheck.RowsAffected != 0 {
		err = errors.New("登录用户名已存在")
		return
	}

	user = models.User{
		LoginName: params.LoginName,
		Name:      params.Name,
		Mobile:    params.Mobile,
		Password:  utils.BcryptMake([]byte(params.Password)),
	}

	err = global.DB.Create(&user).Error
	return
}
