package v1

import (
	"server/src/controller/common"
	"server/src/controller/v1/correlation"
	"server/src/controller/v1/element"
	"server/src/controller/v1/menu"
	"server/src/controller/v1/orifice"
	"server/src/controller/v1/roles"

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
	r.POST("/menu/modify", menu.Update)
	r.POST("/menu/modify/parent", common.Update(tableNameMenu, "parent"))
	r.GET("/menu/list", menu.List)

	// 接口
	r.POST("/interface/add", orifice.Additional)
	r.POST("/interface/delete", common.Delete(tableNameInterface))
	r.POST("/interface/modify", orifice.Update)
	r.POST("/interface/modify/point", common.Update(tableNameInterface, "menu_id"))
	r.GET("/interface/list", orifice.List)

	// 元素
	r.POST("/element/add", element.Additional)
	r.POST("/element/delete", common.Delete(tableNameElement))
	r.POST("/element/modify", element.Update)
	r.POST("/element/modify/point", common.Update(tableNameElement, "menu_id"))
	r.GET("element/list", element.List)

	// 关联表，权限控制
	r.POST("/correlation/add", correlation.Additional)
	r.POST("/correlation/delete", correlation.DeleteCorrelation)

	// 角色
	r.GET("/roles/list", roles.List)

	// 后端需同步接口
	r.GET("/interface/powerlist", orifice.PowerList)
	r.POST("/roles/add", roles.Additional)
	r.POST("/roles/modify", common.Update(tableNameRoles, "role"))
	r.POST("/roles/delete", common.Delete(tableNameRoles))

}
