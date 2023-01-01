package menu

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

	rows1 := spider.MenuList()
	rows2 := spider.MenuPowerList(params.Role)

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
