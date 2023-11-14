package main

import (
	"server/configs"
	v1 "server/src/controller/v1"
	"server/src/middleware"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 接口服务
func server() *gin.Engine {

	app := gin.Default()
	app.Use(middleware.Core)

	base := app.Group(configs.Config.Prefix)
	base.Use(middleware.Recover)
	base.Use(middleware.Logs)
	base.Use(middleware.BodyDispose)

	v1.Route(base.Group("/v1/api"))

	return app
}

// 静态资源
func static() *gin.Engine {
	app := gin.Default()
	app.LoadHTMLGlob("web/*")
	app.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})

	return app
}

func main() {

	// go server().Run(":20020")
	// static().Run(":20030")

	port := ":" + strconv.Itoa(configs.Config.Port)
	server().Run(port)
}
