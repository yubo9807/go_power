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

	// 已存在的角色不允许添加
	rows := spider.Roles.RoleList(params.Role)
	if len(rows) > 0 {
		service.State.ErrorCustom(ctx, "角色'"+params.Role+"'已存在")
		return
	}

	spider.Roles.Additional(params.Role, params.Remark)
	service.State.Success(ctx)
}
