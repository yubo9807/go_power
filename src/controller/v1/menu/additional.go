package menu

import (
	"server/src/service"
	"server/src/spider"

	"github.com/gin-gonic/gin"
)

// 添加菜单
func Additional(ctx *gin.Context) {
	type Params struct {
		Name   string  `form:"name" binding:"required"`
		Title  string  `form:"title" binding:"required"`
		Parent *string `form:"parent"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.ErrorParams()
		return
	}

	// 已存在的菜单不允许添加
	rows := spider.MenuQuery(params.Name, "")
	if len(rows) > 0 {
		service.ErrorCustom("菜单'" + params.Name + "'已存在")
		return
	}

	// 添加菜单
	spider.MenuAdditional(params.Name, params.Title, params.Parent)
	service.Success()
}
