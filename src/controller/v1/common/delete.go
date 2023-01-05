package common

import (
	"server/src/service"
	"server/src/spider"

	"github.com/gin-gonic/gin"
)

// 删除某张表中的某条数据，顺带将关联表中的数据删除
func Delete(tableName string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		type Params struct {
			Id string `form:"id" binding:"required"`
		}
		var params Params
		if err := ctx.ShouldBind(&params); err != nil {
			service.State.ErrorParams()
			return
		}

		spider.Common.Delete(tableName, params.Id)
		spider.Correlation.DeleteCorrelation(tableName, params.Id)
		service.State.Success()
	}
}
