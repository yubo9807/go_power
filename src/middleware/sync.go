package middleware

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
)

var mu sync.Mutex
var wg sync.WaitGroup

// 同步执行中间件
func Sync(ctx *gin.Context) {
	err := ""

	wg.Add(1)
	go func() {
		defer func() {
			if msg := recover(); msg != nil {
				err = fmt.Sprintf("%v", msg)
			}
			mu.Unlock()
			wg.Done()
		}()
		mu.Lock()

		ctx.Next() // 执行下个中间件
	}()
	wg.Wait()

	if err != "" {
		panic(err) // 抛出错误，交给日志中间件处理
	}
}
