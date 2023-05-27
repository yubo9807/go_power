package middleware

import (
	"server/src/service"

	"github.com/gin-gonic/gin"
)

func Authorization(ctx *gin.Context) {

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
