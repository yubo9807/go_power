package orifice

import (
	"server/src/service"
	"server/src/spider"

	"github.com/gin-gonic/gin"
)

func Update(ctx *gin.Context) {
	type Params struct {
		Id     string  `form:"id" binding:"required"`
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

	spider.Interface.Modify(params.Id, params.Method, params.Url, params.Name, params.MenuId)
	service.State.Success()
}
