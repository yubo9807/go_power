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
		service.ErrorParams()
		return
	}

	rows := spider.ElememtQuery(params.Key, "")
	if len(rows) > 0 {
		service.ErrorCustom("元素已存在")
		return
	}

	spider.ElememtAdditional(params.Key, params.Name, params.MenuId)
	service.Success()
}
