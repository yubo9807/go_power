package middleware

import (
	"fmt"
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
		service.State.ErrorUnauthorized(ctx)
		ctx.Abort()
		return
	}

	info, err := service.Jwt.Verify(auth)
	cacheToken := fmt.Sprintf("%v", info["username"])

	// 校验失败，或与缓存中的偷啃不一致
	if err != nil || service.Jwt.StorageGetToken(cacheToken) != auth {
		service.State.ErrorTokenFailure(ctx)
		ctx.Abort()
		return
	}

}
