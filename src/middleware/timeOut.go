package middleware

import (
	"context"
	"server/configs"
	"time"

	"github.com/gin-gonic/gin"
)

func Timeout(ctx *gin.Context) {
	context, cancel := context.WithTimeout(ctx.Request.Context(), time.Duration(configs.Config.TimeOut)*time.Second)
	defer cancel()

	ctx.Request = ctx.Request.WithContext(context)

	ch := make(chan struct{})
	go func() {
		ctx.Next()
		close(ch)
	}()

	select {
	case <-ch:
		ctx.Next()
	case <-context.Done():
		ctx.String(504, "connect timeout")
	}
}
