package correlation

import (
	"server/src/service"
	"server/src/spider"

	"github.com/gin-gonic/gin"
)

func Additional(ctx *gin.Context) {
	type Params struct {
		RoleId    string `form:"roleId" binding:"required"`
		TableId   string `form:"tableId" binding:"required"`
		TableType string `form:"tableType" binding:"required"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.ErrorParams()
		return
	}

	// 已存在的关联不允许添加
	rows := spider.CorrelationQuery(params.RoleId, params.TableId, params.TableType)
	if len(rows) > 0 {
		service.ErrorCustom("关联信息已存在")
		return
	}

	spider.CorrelationAdditional(params.RoleId, params.TableId, params.TableType)
	service.Success()
}
