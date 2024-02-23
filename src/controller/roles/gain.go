package roles

import (
	"server/src/service"
	"server/src/spider"

	"github.com/gin-gonic/gin"
)

func List(ctx *gin.Context) {
	type Params struct {
		Role string `form:"role"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.State.ErrorParams(ctx)
		return
	}

	rows := spider.Roles.RoleList(params.Role)
	service.State.SuccessData(ctx, rows)
}
