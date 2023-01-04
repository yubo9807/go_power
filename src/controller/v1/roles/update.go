package roles

import (
	"server/src/service"
	"server/src/spider"

	"github.com/gin-gonic/gin"
)

func Update(ctx *gin.Context) {
	type Params struct {
		Id     string `form:"id" binding:"required"`
		Role   string `form:"role" binding:"required"`
		Remark string `form:"remark"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.ErrorParams()
		return
	}

	spider.CommonUpdate("element", params.Id, "role", params.Role)
	service.Success()
}
