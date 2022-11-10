/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"monaPanel/bootstrap"
	"monaPanel/global"
)

func main() {

	// 初始化配置文件，读取应用配置
	bootstrap.InitConfig()

	// 初始化log模块
	global.Log = bootstrap.InitLog()
	global.Log.Info("Log system start ok.")

	// 初始化 gorm 数据库
	global.DB = bootstrap.InitDatabase()
	defer func() {
		if global.DB != nil {
			db, _ := global.DB.DB()
			db.Close()
		}
	}()

	// 初始化Redis
	global.Redis = bootstrap.InitializeRedis()

	// 初始化验证器
	bootstrap.InitializeValidator()

	// 启动 go-gin web服务器
	bootstrap.RunServer()
}
