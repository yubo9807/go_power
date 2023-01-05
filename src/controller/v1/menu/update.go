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
		service.State.ErrorParams()
		return
	}

	bool := detectionStructure(params.Id, params.Parent, make([]string, 0))
	if params.Parent != nil && params.Id == *params.Parent || bool {
		service.State.ErrorCustom("父级菜单指向错误")
		return
	}

	spider.Menu.Modify(params.Id, params.Name, params.Title, params.Parent)
	service.State.Success()
}

// 结构检查，防止出现循环指向
// 检测到返回 true，没有则返回 false
func detectionStructure(id string, parent *string, collect []string) bool {
	if parent == nil {
		return false
	}

	check := false
	rows := spider.Menu.StructureQuery(parent)
	collect = append(collect, rows[0].Id)
	for i := 0; i < len(collect); i++ {
		if collect[i] == id {
			check = true
			break
		}
	}
	if rows[0].Parent != nil {
		check = detectionStructure(id, rows[0].Parent, collect)
	}
	return check
}
