package main

import (
	"zxs-gin/bootstrap"
	"zxs-gin/global"
)

func main() {
	// 初始化配置
	bootstrap.InitializeConfig()
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("init success")

	// 初始化数据库
	global.App.DB = bootstrap.InitializeDB()
	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()

	// 启动服务
	bootstrap.RunServer()
}
