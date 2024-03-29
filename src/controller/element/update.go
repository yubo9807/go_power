package element

import (
	"server/src/service"
	"server/src/spider"

	"github.com/gin-gonic/gin"
)

func Update(ctx *gin.Context) {
	type Params struct {
		Id     string `form:"id" binding:"required"`
		Key    string `form:"key" binding:"required"`
		Name   string `form:"name" binding:"required"`
		MenuId string `form:"menuId"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.State.ErrorParams(ctx)
		return
	}

	spider.Elememt.Modify(params.Id, params.Key, params.Name, params.MenuId)
	service.State.Success(ctx)
}
