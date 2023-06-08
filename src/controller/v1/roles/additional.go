package roles

import (
	"server/src/service"
	"server/src/spider"

	"github.com/gin-gonic/gin"
)

func Additional(ctx *gin.Context) {
	type Params struct {
		Role   string `form:"role" binding:"required"`
		Remark string `form:"remark"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.State.ErrorParams(ctx)
		return
	}

	spider.Roles.Additional(params.Role)
	service.State.Success(ctx)
}
