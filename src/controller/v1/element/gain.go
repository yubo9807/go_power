package element

import (
	"server/src/service"
	"server/src/spider"

	"github.com/gin-gonic/gin"
)

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

	rows1 := spider.ElememtList(params.Point)
	rows2 := spider.ElememtPowerList(params.Role, params.Point)

	for i := 0; i < len(rows1); i++ {
		for j := 0; j < len(rows2); j++ {
			if rows2[j].Key == rows1[i].Key {
				rows1[i].Selected = true
				rows1[i].CorrelationId = rows2[j].CorrelationId
				rows1[i].RoleId = rows2[j].RoleId
			}
		}
	}

	data := []spider.Elememt{}
	data = append(data, rows1...)
	service.SuccessData(data)
}
