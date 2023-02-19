package middleware

import (
	"server/configs"

	"github.com/gin-gonic/gin"
)

func Core(ctx *gin.Context) {
	origin := "http://hpyyb.cn"
	if configs.Env.DEVELOPMENT {
		origin = "*"
	}

	ctx.Header("Access-Control-Allow-Origin", origin)
	ctx.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type")

	if ctx.Request.Method == "OPTIONS" {
		ctx.String(200, "ok")
	}
}
