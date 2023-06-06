package menu

import (
	"server/src/service"
	"server/src/spider"
	"sync"

	"github.com/gin-gonic/gin"
)

var mu sync.Mutex
var wg sync.WaitGroup

// 添加菜单
func Additional(ctx *gin.Context) {
	type Params struct {
		Name   string  `form:"name" binding:"required"`
		Title  string  `form:"title" binding:"required"`
		Hidden bool    `form:"hidden"`
		Parent *string `form:"parent"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.State.ErrorParams()
		return
	}

	wg.Add(1)
	go func() {
		defer func() { mu.Unlock(); wg.Done() }()
		mu.Lock()

		// 已存在的菜单不允许添加
		rows := spider.Menu.Query(params.Name, "")
		if len(rows) > 0 {
			service.State.ErrorCustom("菜单'" + params.Name + "'已存在")
			return
		}
		// 添加菜单
		spider.Menu.Additional(params.Name, params.Title, params.Hidden, params.Parent)
		service.State.Success()
	}()
	wg.Wait()
}
