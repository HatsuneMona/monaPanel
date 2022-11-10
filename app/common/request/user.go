package request

type Register struct {
	LoginName string `json:"loginName" form:"loginName" binding:"required"`
	Name      string `json:"name" form:"name" binding:"required"`
	Mobile    string `json:"mobile" form:"mobile" binding:"mobile"`
	Password  string `json:"password" form:"name" binding:"required"`
}

type Login struct {
	LoginName string `json:"loginName" form:"loginName" binding:"required"`
	Password  string `json:"password" form:"password" binding:"required"`
}

func (r Register) GetMessages() ValidatorMassages {
	return ValidatorMassages{
		"loginName.required": "登录用户名不能为空",
		"name.required":      "用户昵称不能为空",
		"mobile.mobile":      "手机号码不符合规范",
		"password.required":  "密码不能为空",
	}
}

func (l Login) GetMessages() ValidatorMassages {
	return ValidatorMassages{
		"loginName.required": "登录用户名不能为空",
		"password.required":  "密码不能为空",
	}
}
