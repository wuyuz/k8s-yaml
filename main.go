package main

import (
	. "k8s_yaml/app"
	"k8s_yaml/config"
	"k8s_yaml/lib"
	"k8s_yaml/routes"
)

func main()  {
	// 日志wewe初始化
	Log = lib.LogInit()

	// 命令解析parse

	Serve()
}

func Serve() {
	Boot()
	// 路由test加载
	routes.LoadRouter()

	// 启动监听服务
	if err := App.Listen(":8090"); err != nil {
		Log.Warn(err.Error())
		return
	}

}

// 初始化引导
func Boot() {
	config.BootApp()
}