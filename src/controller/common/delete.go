package common

import (
	"server/src/service"
	"server/src/spider"

	"github.com/gin-gonic/gin"
)

// 删除某张表中的某条数据
func Delete(tableName string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		type Params struct {
			id string `form:"id" binding:"required"`
		}

		var params Params
		if err := ctx.ShouldBind(&params); err != nil {
			service.ErrorParams()
			return
		}

		spider.CommonDelete(tableName, params.id)
		service.Success()
	}
}
