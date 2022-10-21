package bootstrap

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"monaPanel/global"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)
import "monaPanel/routers"

func setupRouter() *gin.Engine {
	router := gin.Default()

	// 静态资源
	router.StaticFile("/", "./static/dist/index.html")
	router.Static("/assets", "./static/dist/assets")
	router.StaticFile("/favicon.ico", "./static/dist/favicon.ico")

	// api 路由
	apiGroup := router.Group("/api")
	routers.SetApiGroupRoutes(apiGroup)
	return router
}

// RunServer 启动服务器
func RunServer() {
	r := setupRouter()
	r.Run(global.Config.Server.Listen + ":" + global.Config.Server.Port)

	srv := &http.Server{
		Addr:    global.Config.Server.Listen + ":" + global.Config.Server.Port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Log.Fatal(fmt.Sprintf("listen err: %s\n", err))
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	global.Log.Info("shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.Log.Fatal(fmt.Sprintf("Server Shutdown Err: %s\n", err))
	}
	global.Log.Info("Server exiting")
}
