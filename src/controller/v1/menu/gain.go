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
				rows1[i].Correlation_id = rows2[j].Correlation_id
				rows1[i].Role_id = rows2[j].Role_id
			}
		}
	}

	data := []spider.Menu{}
	data = append(data, rows1...)
	service.SuccessData(data)
}
