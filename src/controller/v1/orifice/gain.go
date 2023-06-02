package orifice

import (
	"server/src/service"
	"server/src/spider"

	"github.com/gin-gonic/gin"
)

// 获取接口数据（按模块）
func List(ctx *gin.Context) {
	type Params struct {
		RoleId string `form:"roleId" binding:"required"`
		MenuId string `form:"menuId"`
		Url    string `form:"url"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.State.ErrorParams()
		return
	}

	rows1 := spider.Interface.List(params.MenuId, params.Url)
	rows2 := spider.Interface.PowerListModule(params.RoleId, params.MenuId)

	for i := 0; i < len(rows1); i++ {
		for j := 0; j < len(rows2); j++ {
			if rows2[j].Url == rows1[i].Url && rows2[j].Method == rows1[i].Method {
				rows1[i].Selected = true
				rows1[i].CorrelationId = rows2[j].CorrelationId
				rows1[i].RoleId = rows2[j].RoleId
			}
		}
	}

	data := []spider.InterfaceColumn{}
	data = append(data, rows1...)
	service.State.SuccessData(data)
}

// 获取具有权限的所有接口
func Authority(ctx *gin.Context) {
	type Params struct {
		RoleId string `form:"roleId" binding:"required"`
		Method string `form:"method" binding:"required"`
		Url    string `form:"url" binding:"required"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.State.ErrorParams()
		return
	}

	rows := spider.Interface.PowerList(params.RoleId, params.Method, params.Url)
	data := []spider.InterfaceColumn{}
	data = append(data, rows...)
	if len(data) > 0 {
		service.State.SuccessData(data[0])
	} else {
		service.State.ErrorCustom("permission error")
	}
}
