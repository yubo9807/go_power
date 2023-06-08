package roles

import (
	"server/src/service"
	"server/src/spider"

	"github.com/gin-gonic/gin"
)

func List(ctx *gin.Context) {
	rows := spider.Roles.RoleList()
	service.State.SuccessData(ctx, rows)
}
