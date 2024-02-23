package router

import (
	"server/configs"
	"server/src/controller/common"
	"server/src/controller/correlation"
	"server/src/controller/element"
	"server/src/controller/menu"
	"server/src/controller/orifice"
	"server/src/controller/roles"
	"server/src/middleware"

	"github.com/gin-gonic/gin"
)

func V1(r *gin.RouterGroup) {

	// 授权验证
	r.Use(middleware.Authorization)

	// 菜单
	r.POST("/menu/add", middleware.Sync, menu.Additional)
	r.POST("/menu/delete", menu.Delete)
	r.POST("/menu/modify", menu.Update)
	r.GET("/menu/list", menu.List)
	r.POST("/menu/sort", menu.UpdateSort)

	// 接口
	r.POST("/interface/add", middleware.Sync, orifice.Additional)
	r.POST("/interface/delete", common.Delete(configs.Table_Interface))
	r.POST("/interface/modify", orifice.Update)
	r.GET("/interface/list", orifice.List)

	// 元素
	r.POST("/element/add", middleware.Sync, element.Additional)
	r.POST("/element/delete", common.Delete(configs.Table_Element))
	r.POST("/element/modify", element.Update)
	r.GET("/element/list", element.List)
	r.GET("/element/authority", element.Authority)

	// 关联表，权限控制
	r.POST("/correlation/synchronization", correlation.Synchronization)

	// 后端拿到该角色有权限的接口，做中间件拦截
	r.GET("/interface/authority", orifice.Authority)

	// 业务系统进行角色同步
	r.GET("/roles/list", roles.List)
	r.POST("/roles/add", middleware.Sync, roles.Additional)
	r.POST("/roles/modify", roles.Update)
	r.POST("/roles/delete", roles.Delete)

}
