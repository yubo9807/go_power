package menu

import (
	"server/src/service"
	"server/src/spider"

	"github.com/gin-gonic/gin"
)

// 获取所有菜单数据
func List(ctx *gin.Context) {
	type Params struct {
		RoleId string `form:"roleId" binding:"required"`
		Title  string `form:"title"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.ErrorParams()
		return
	}

	rows1 := spider.MenuList(params.Title)
	rows2 := spider.MenuPowerList(params.RoleId)

	for i := 0; i < len(rows1); i++ {
		for j := 0; j < len(rows2); j++ {
			if rows2[j].Name == rows1[i].Name {
				rows1[i].Selected = true
				rows1[i].CorrelationId = rows2[j].CorrelationId
				rows1[i].RoleId = rows2[j].RoleId
			}
		}
	}

	data := []spider.Menu{}
	data = append(data, rows1...)
	service.SuccessData(data)
}
