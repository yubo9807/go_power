package orifice

import (
	"server/src/service"
	"server/src/spider"

	"github.com/gin-gonic/gin"
)

// 添加角色
func Additional(ctx *gin.Context) {
	type Params struct {
		Method string  `form:"method" binding:"required"`
		Url    string  `form:"url" binding:"required"`
		Name   string  `form:"name" binding:"required"`
		MenuId *string `form:"menuId"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.State.ErrorParams()
		return
	}

	rows := spider.Interface.Query(params.Method, params.Url)
	if len(rows) > 0 {
		service.State.ErrorCustom("接口已存在")
		return
	}

	spider.Interface.Additional(params.Method, params.Url, params.Name, params.MenuId)
	service.State.Success()
}
