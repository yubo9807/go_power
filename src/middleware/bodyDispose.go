package middleware

import (
	"server/src/service"

	"github.com/gin-gonic/gin"
)

func BodyDispose(ctx *gin.Context) {
	service.State.InitState(ctx)

	go func() {
		sql := <-service.ChSql
		service.State.RecordSql(ctx, sql[0], sql[1])
	}()

	ctx.Next()

	// 如果已经返回了结果，不对数据进行包装
	if ctx.Writer.Written() {
		return
	}

	service.State.Result(ctx)

}
