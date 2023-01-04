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
		service.ErrorParams()
		return
	}

	rows1 := spider.InterfaceList(params.MenuId, params.Url)
	rows2 := spider.InterfacePowerList(params.RoleId, params.MenuId)

	for i := 0; i < len(rows1); i++ {
		for j := 0; j < len(rows2); j++ {
			if rows2[j].Url == rows1[i].Url && rows2[j].Method == rows1[i].Method {
				rows1[i].Selected = true
				rows1[i].CorrelationId = rows2[j].CorrelationId
				rows1[i].RoleId = rows2[j].RoleId
			}
		}
	}

	data := []spider.Interface{}
	data = append(data, rows1...)
	service.SuccessData(data)
}

// 获取具有权限的所有接口
func PowerListAll(ctx *gin.Context) {
	type Params struct {
		RoleId string `form:"roleId" binding:"required"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.ErrorParams()
		return
	}

	rows := spider.InterfacePowerListAll(params.RoleId)
	data := []spider.Interface{}
	data = append(data, rows...)
	service.SuccessData(data)
}
