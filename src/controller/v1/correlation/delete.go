package correlation

import (
	"server/src/service"
	"server/src/spider"

	"github.com/gin-gonic/gin"
)

// 删除关联信息
func DeleteCorrelation(ctx *gin.Context) {
	type Params struct {
		TableId string `form:"tableId" binding:"required"`
	}

	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.ErrorParams()
		return
	}

	spider.CorrelationDeleteCorrelation(params.TableId)
	service.Success()
}
