package middleware

import (
	"server/src/service"

	"github.com/gin-gonic/gin"
)

func BodyDispose(ctx *gin.Context) {

	service.State.InitState(ctx)

	ctx.Next()
	service.State.RecordSql(ctx, service.SqlStrs)
	service.SqlStrs = service.SqlStrs[:0]

	// 如果已经返回了结果，不对数据进行包装
	if ctx.Writer.Written() {
		return
	}

	service.State.Result(ctx)

}
