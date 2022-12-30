package orifice

import (
	"server/src/service"
	"server/src/spider"

	"github.com/gin-gonic/gin"
)

// 获取所有菜单数据
func List(ctx *gin.Context) {
	type Params struct {
		Role string `form:"role" binding:"required"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.ErrorParams()
		return
	}

	rows1 := spider.InterfaceList()
	rows2 := spider.InterfacePowerList(params.Role)

	for i := 0; i < len(rows1); i++ {
		for j := 0; j < len(rows2); j++ {
			if rows2[j].Url == rows1[i].Url {
				rows1[i].Selected = true
			}
		}
	}

	service.SuccessData(rows1)
}

// 查询菜单
func Query(ctx *gin.Context) {
	type Params struct {
		Name  string `form:"name"`
		Title string `form:"title"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.ErrorParams()
		return
	}

	rows := spider.MenuQuery(params.Name, params.Title)

	service.SuccessData(rows)
}
