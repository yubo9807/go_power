package menu

import (
	"server/src/service"
	"server/src/spider"

	"github.com/gin-gonic/gin"
)

// 更新基本信息
func Update(ctx *gin.Context) {
	type Params struct {
		Id     string  `form:"id" binding:"required"`
		Name   string  `form:"name" binding:"required"`
		Title  string  `form:"title"`
		Parent *string `form:"parent"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.ErrorParams()
		return
	}

	spider.MenuModify(params.Id, params.Name, params.Title, params.Parent)
	service.Success()
}
