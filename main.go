/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"monaPanel/bootstrap"
	"monaPanel/global"
	"net/http"
)

func main() {
	bootstrap.InitConfig()
	global.Log = bootstrap.InitLog()

	global.Log.Info("Log system start ok.")

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	serverConf := global.Config.Server

	r.Run(fmt.Sprintf("%s:%s", serverConf.Listen, serverConf.Port))
}
