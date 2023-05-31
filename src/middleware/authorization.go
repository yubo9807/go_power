package middleware

import (
	"server/configs"
	"server/src/service"

	"github.com/gin-gonic/gin"
)

func Authorization(ctx *gin.Context) {

	if !configs.Config.Certification {
		ctx.Next()
		return
	}

	auth := ctx.GetHeader("Authorization")
	if auth == "" {
		service.State.ErrorUnauthorized()
		ctx.Abort()
		return
	}

	_, err := service.Jwt.Verify(auth)
	if err != nil {
		service.State.ErrorTokenFailure()
		ctx.Abort()
		return
	}

	// ctx.Next()

}
