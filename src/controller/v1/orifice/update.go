package orifice

import (
	"server/src/service"
	"server/src/spider"

	"github.com/gin-gonic/gin"
)

func Update(ctx *gin.Context) {
	type Params struct {
		Id     string `form:"id" binding:"required"`
		Method string `form:"method" binding:"required"`
		Url    string `form:"url" binding:"required"`
		Name   string `form:"name" binding:"required"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.ErrorParams()
		return
	}

	rows := spider.InterfaceQuery(params.Method, params.Url)
	if len(rows) > 0 {
		service.ErrorCustom("接口已存在")
		return
	}

	spider.InterfaceModify(params.Id, params.Method, params.Url, params.Name)
}
