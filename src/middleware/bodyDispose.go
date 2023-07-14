package middleware

import (
	"server/src/service"
	"sync"

	"github.com/gin-gonic/gin"
)

var mu sync.Mutex
var wg sync.WaitGroup

func BodyDispose(ctx *gin.Context) {

	// wg.Add(1)
	// go func() {
	// 	defer func() { mu.Unlock(); wg.Done() }()
	// 	mu.Lock()

	service.State.InitState(ctx)

	ctx.Next()
	service.State.RecordSql(ctx, service.SqlStrs)
	service.SqlStrs = service.SqlStrs[:0]

	// 如果已经返回了结果，不对数据进行包装
	if ctx.Writer.Written() {
		return
	}

	service.State.Result(ctx)
	// }()
	// wg.Wait()

}
