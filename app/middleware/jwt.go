package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"monaPanel/app/common/response"
	"monaPanel/app/service"
	"monaPanel/global"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Authorization")
		if tokenStr == "" {
			response.ClaimsTokenFail(c)
			c.Abort()
			return
		}

		// Token 解析校验
		token, err := jwt.ParseWithClaims(
			tokenStr, &service.AdminUserClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(global.Config.Jwt.Secret), nil
			},
		)
		// 如果验证不通过（token无法解码），或token在黑名单中
		// 则返回token错误
		if err != nil || service.JwtService.IsInBlacklist(tokenStr) {
			response.ClaimsTokenFail(c)
			c.Abort()
			return
		}

		claims := token.Claims.(*service.AdminUserClaims)
		// Token 发布者校验
		c.Set("token", token)
		c.Set("userId", claims.UserId)
	}
}
