package v1

import (
	"server/src/controller/common"
	"server/src/controller/v1/correlation"
	"server/src/controller/v1/menu"
	"server/src/controller/v1/orifice"

	"github.com/gin-gonic/gin"
)

const (
	tableNameMenu      = "menu"
	tableNameInterface = "interface"
	tableNameElement   = "element"
	tableNameRoles     = "roles"
)

func Route(r *gin.RouterGroup) {

	// 菜单
	r.POST("/menu/add", menu.Additional)
	r.POST("/menu/delete", common.Delete(tableNameMenu))
	r.POST("/menu/update", menu.Update)
	r.POST("/menu/update/parent", common.Update(tableNameMenu, "parent"))
	r.GET("/menu/list", menu.List)

	// 接口
	r.POST("/interface/delete", common.Delete(tableNameInterface))
	r.POST("/interface/update/menu", common.Update(tableNameInterface, "point"))
	r.GET("/interface/list", orifice.List)

	// 元素
	r.POST("/element/delete", common.Delete(tableNameElement))
	r.POST("/element/update/menu", common.Update(tableNameElement, "point"))

	// 角色
	r.POST("/roles/delete", common.Delete(tableNameRoles))
	r.POST("/roles/update", common.Update(tableNameRoles, "role"))

	// 关联表
	r.POST("/correlation/add", correlation.Additional)
	r.POST("/correlation/delete", correlation.DeleteCorrelation)

}
