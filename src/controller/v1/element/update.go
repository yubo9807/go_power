package element

import (
	"server/src/service"
	"server/src/spider"

	"github.com/gin-gonic/gin"
)

func Update(ctx *gin.Context) {
	type Params struct {
		Id   string `form:"id" binding:"required"`
		Key  string `form:"key" binding:"required"`
		Name string `form:"name" binding:"required"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.ErrorParams()
		return
	}

	rows := spider.ElememtQuery(params.Key, "")
	if len(rows) > 0 {
		service.ErrorCustom("元素已存在")
		return
	}

	spider.ElememtModify(params.Id, params.Key, params.Name)
}
