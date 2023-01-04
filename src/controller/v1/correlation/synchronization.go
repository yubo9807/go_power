package correlation

import (
	"server/src/service"
	"server/src/spider"

	"github.com/gin-gonic/gin"
)

// 同步权限
func Synchronization(ctx *gin.Context) {
	type Params struct {
		RoleId        string   `form:"roleId" binding:"required"`
		ContactIdList []string `form:"contactIdList" binding:"required"`
		DeleteIdList  []string `form:"deleteIdList" binding:"required"`
		TableType     string   `form:"tableType" binding:"required"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.ErrorParams()
		return
	}

	newContactIdList := make([]string, 0)
	rows := spider.CorrelationTableTypeQuery(params.RoleId, params.TableType)

	for i := 0; i < len(params.ContactIdList); i++ {
		exist := false
		for j := 0; j < len(rows); j++ {
			exist = params.ContactIdList[i] == rows[j].TableId
			if exist {
				break
			}
		}
		if !exist { // 已经不存在
			newContactIdList = append(newContactIdList, params.ContactIdList[i])
		}
	}

	spider.CorrelationBatchAdditional(params.TableType, params.RoleId, newContactIdList, params.DeleteIdList)
	service.Success()

}
