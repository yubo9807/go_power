package orifice

import (
	"server/src/service"
	"server/src/spider"

	"github.com/gin-gonic/gin"
)

// 获取所有菜单数据
func List(ctx *gin.Context) {
	type Params struct {
		Role  string `form:"role" binding:"required"`
		Point string `form:"point"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.ErrorParams()
		return
	}

	rows1 := spider.InterfaceList(params.Point)
	rows2 := spider.InterfacePowerList(params.Role, params.Point)

	for i := 0; i < len(rows1); i++ {
		for j := 0; j < len(rows2); j++ {
			if rows2[j].Url == rows1[i].Url {
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
func PowerList(ctx *gin.Context) {
	type Params struct {
		Role  string `form:"role" binding:"required"`
		Point string `form:"point"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.ErrorParams()
		return
	}

	rows := spider.InterfacePowerList(params.Role, params.Point)
	data := []spider.Interface{}
	data = append(data, rows...)
	service.SuccessData(data)
}
