package orifice

import (
	"server/configs"
	"server/src/service"
	"server/src/spider"

	"github.com/gin-gonic/gin"
)

// 删除某张表中的某条数据，顺带将关联表中的数据删除
func Delete(ctx *gin.Context) {
	type Params struct {
		Id string `form:"id" binding:"required"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.State.ErrorParams(ctx)
		return
	}

	spider.Correlation.DeleteCorrelation("interface", params.Id)
	spider.Common.Delete(configs.Table_Interface, params.Id)
	service.State.Success(ctx)
}
