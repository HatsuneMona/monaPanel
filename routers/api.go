package routers

import (
	"github.com/gin-gonic/gin"
	"monaPanel/controllers"
	"net/http"
)

func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.POST("/auth/register", controllers.Register)
}
