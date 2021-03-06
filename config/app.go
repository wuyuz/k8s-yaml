package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/django"
	. "k8s_yaml/app"
	"net/http"
)

func BootApp() {
	// 加载环境变量

	// 初始化模板渲染
	TemplateEngine = django.NewFileSystem(http.Dir("./resource/views"),".html")


	// 初始化App
	App = fiber.New(fiber.Config{
		ServerHeader: "k8s-yaml",
		Views: TemplateEngine,
		// BodyLimit:
	})

	App.Use(recover.New()) // 报错拦截
	App.Use(pprof.New())  // 性能分析

	// 静态文件
	App.Static("/assets","resource/assets",fiber.Static{
		Compress: true,
	})
}
