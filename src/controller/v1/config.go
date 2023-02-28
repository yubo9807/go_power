package v1

import (
	"server/src/controller/v1/common"
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
	r.POST("/menu/delete", menu.Delete)
	r.POST("/menu/modify", menu.Update)
	r.GET("/menu/list", menu.List)

	// 接口
	r.POST("/interface/add", orifice.Additional)
	r.POST("/interface/delete", common.Delete(tableNameInterface))
	r.POST("/interface/modify", orifice.Update)
	r.GET("/interface/list", orifice.List)

	// 元素
	r.POST("/element/add", element.Additional)
	r.POST("/element/delete", common.Delete(tableNameElement))
	r.POST("/element/modify", element.Update)
	r.GET("/element/list", element.List)

	// 关联表，权限控制
	r.POST("/correlation/synchronization", correlation.Synchronization)

	// 后端拿到该角色有权限的接口，做中间件拦截
	r.GET("/interface/powerlist", orifice.PowerListAll)

	// 业务系统进行角色同步
	r.GET("/roles/list", roles.List)
	r.POST("/roles/add", roles.Additional)
	r.POST("/roles/modify", roles.Update)
	r.POST("/roles/delete", roles.Delete)

}
