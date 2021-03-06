package routes

import (
	"github.com/gofiber/fiber/v2"
	. "k8s_yaml/app"
	"k8s_yaml/contraller"
)

func webRoute () {
	var (
		web fiber.Router
	)
	web = App.Group("")
	LoadingRoutes(web)
}

func LoadingRoutes(app fiber.Router) {
	// 子级中间件
	app.Get("/index", contraller.Landing)
	app.Get("/load",contraller.Loading)
}
