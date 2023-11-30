package middleware

import (
	"context"
	"server/configs"
	"server/src/service"
	"time"

	"github.com/gin-gonic/gin"
)

func Timeout(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(configs.Config.TimeOut)*time.Second)
	defer cancel()

	c.Request = c.Request.WithContext(ctx)

	ch := make(chan struct{})
	go func() {
		c.Next()
		close(ch)
	}()

	select {
	case <-ch:
		c.Next()
	case <-ctx.Done():
		service.State.ErrorRequestTimeout(c)
	}
}
