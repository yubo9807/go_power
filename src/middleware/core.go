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
}
