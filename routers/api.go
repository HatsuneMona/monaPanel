package routers

import (
	"github.com/gin-gonic/gin"
	"monaPanel/controllers"
	"monaPanel/middleware"
	"net/http"
)

func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.GET(
		"/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		},
	)

	router.POST("/auth/register", controllers.Register)
	router.POST("/auth/login", controllers.Login)

	authRouter := router.Group("").Use(middleware.JwtAuth())
	{
		authRouter.POST("/auth/getUserInfo", controllers.GetUserInfo)
	}
}
