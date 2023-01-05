package element

import (
	"server/src/service"
	"server/src/spider"

	"github.com/gin-gonic/gin"
)

func Additional(ctx *gin.Context) {
	type Params struct {
		Key    string  `form:"key" binding:"required"`
		Name   string  `form:"name" binding:"required"`
		MenuId *string `form:"menuId"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.State.ErrorParams()
		return
	}

	rows := spider.Elememt.Query(params.Key, "")
	if len(rows) > 0 {
		service.State.ErrorCustom("元素已存在")
		return
	}

	spider.Elememt.Additional(params.Key, params.Name, params.MenuId)
	service.State.Success()
}
