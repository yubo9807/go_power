package roles

import (
	"server/configs"
	"server/src/service"
	"server/src/spider"

	"github.com/gin-gonic/gin"
)

func Delete(ctx *gin.Context) {
	type Params struct {
		Id string `form:"id" binding:"required"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.State.ErrorParams(ctx)
		return
	}

	spider.Correlation.DeleteCorrelationRoles(params.Id)
	spider.Common.Delete(configs.Table_Roles, params.Id)
	service.State.Success(ctx)
}
