package common

import (
	"server/src/service"
	"server/src/spider"

	"github.com/gin-gonic/gin"
)

// 修改某张表的某一个字段
func Update(tableName, key string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		type Params struct {
			Id    string `form:"id" binding:"required"`
			Value string `form:"value" binding:"required"`
		}

		var params Params
		if err := ctx.ShouldBind(&params); err != nil {
			service.ErrorParams()
			return
		}

		spider.CommonUpdate(tableName, params.Id, key, params.Value)
		service.Success()
	}
}
