package middleware

import (
	"server/configs"
	"server/src/service"

	"github.com/gin-gonic/gin"
)

func Authorization(ctx *gin.Context) {

	auth := ctx.GetHeader("Open-Id")
	if auth != configs.Config.OpenId {
		service.State.ErrorUnauthorized(ctx)
		ctx.Abort()
		return
	}

}
