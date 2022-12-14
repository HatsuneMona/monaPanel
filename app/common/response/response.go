package response

import (
	"github.com/gin-gonic/gin"
	"monaPanel/global"
	"net/http"
)

type Response struct {
	Errno  int         `json:"errno"`
	Errmsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

// Success 响应成功
func Success(c *gin.Context, data interface{}) {
	c.JSON(
		http.StatusOK, Response{
			Errno:  0,
			Errmsg: "ok",
			Data:   data,
		},
	)
}

// Fail 响应失败 ErrorCode 不为 0 表示失败
func Fail(c *gin.Context, errorCode int, msg string) {
	c.JSON(
		http.StatusOK, Response{
			Errno:  errorCode,
			Errmsg: msg,
			Data:   nil,
		},
	)
}

// FailByError 失败响应 返回自定义错误的错误码、错误信息
func FailByError(c *gin.Context, error global.CustomError) {
	Fail(c, error.ErrorCode, error.ErrorMsg)
}

// ValidateFail 请求参数验证失败
func ValidateFail(c *gin.Context, msg string) {
	Fail(c, global.Errors.ValidateError.ErrorCode, msg)
}

// ServiceFail 调用service失败
func ServiceFail(c *gin.Context, msg string) {
	Fail(c, global.Errors.ServiceError.ErrorCode, msg)
}

// ClaimsTokenFail jwt鉴权失败
func ClaimsTokenFail(c *gin.Context) {
	Fail(c, global.Errors.ClaimsTokenError.ErrorCode, global.Errors.ClaimsTokenError.ErrorMsg)
}
