package models

type TestGorm struct {
	ID
	Name     string `json:"name" gorm:"not_null;comment:用户名"`
	Mobile   string `json:"mobile" gorm:"not_null;index;comment:手机号码"`
	Password string `json:"password" gorm:"not_null;default:'';comment:登录密码"`
	CommonTime
	DeleteTime
}
