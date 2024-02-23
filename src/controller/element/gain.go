package element

import (
	"server/src/service"
	"server/src/spider"

	"github.com/gin-gonic/gin"
)

// 获取元素数据（按模块）
func List(ctx *gin.Context) {
	type Params struct {
		RoleId string `form:"roleId" binding:"required"`
		MenuId string `form:"menuId"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.State.ErrorParams(ctx)
		return
	}

	rows1 := spider.Elememt.List(params.MenuId)
	rows2 := spider.Elememt.PowerList(params.RoleId, params.MenuId)

	for i := 0; i < len(rows1); i++ {
		for j := 0; j < len(rows2); j++ {
			if rows2[j].Key == rows1[i].Key {
				rows1[i].Selected = true
				rows1[i].CorrelationId = rows2[j].CorrelationId
				rows1[i].RoleId = rows2[j].RoleId
			}
		}
	}

	data := []spider.ElememtColumn{}
	data = append(data, rows1...)
	service.State.SuccessData(ctx, data)
}

func Authority(ctx *gin.Context) {
	type Params struct {
		RoleId string `form:"roleId" binding:"required"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.State.ErrorParams(ctx)
		return
	}

	rows := spider.Elememt.PowerList2(params.RoleId)

	data := []spider.ElememtColumn{}
	data = append(data, rows...)
	service.State.SuccessData(ctx, data)
}
