package v1

import (
	"server/src/controller/common"
	"server/src/controller/v1/menu"
	"server/src/controller/v1/orifice"

	"github.com/gin-gonic/gin"
)

const (
	tableNameMenu      = "menu"
	tableNameInterface = "interface"
	tableNameButton    = "buttom"
)

func Route(r *gin.RouterGroup) {
	// 菜单
	r.POST("/menu/delete", common.Delete(tableNameMenu))
	r.POST("/menu/update/parent", common.Update(tableNameMenu, "parent"))
	r.GET("/menu/list", menu.List)
	r.POST("/menu/add", menu.Additional)
	r.GET("/menu/query", menu.Query)
	r.POST("/menu/update", menu.Update)

	// 接口
	r.POST("/interface/delete", common.Delete(tableNameInterface))
	r.POST("/interface/update/menu", common.Update(tableNameInterface, tableNameMenu))
	r.GET("/interface/list", orifice.List)

	// btn
	r.POST("/button/delete", common.Delete(tableNameButton))
	r.POST("/button/update/menu", common.Update(tableNameButton, tableNameMenu))
}
