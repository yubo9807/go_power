package correlation

import (
	"server/src/service"
	"server/src/spider"

	"github.com/gin-gonic/gin"
)

// 同步权限
func Synchronization(ctx *gin.Context) {
	type Params struct {
		RoleId      string   `form:"roleId" binding:"required"`
		TableIdList []string `form:"tableIdList" binding:"required"`
		TableType   string   `form:"tableType" binding:"required"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.ErrorParams()
		return
	}

	newTableIdList := make([]string, 0)
	rows := spider.CorrelationTableTypeQuery(params.RoleId, params.TableType)
	for i := 0; i < len(params.TableIdList); i++ {
		exist := false
		for j := 0; j < len(rows); j++ {
			exist = params.TableIdList[i] == rows[j].RoleId && params.TableIdList[i] == rows[j].TableId
			if exist {
				break
			}
		}
		if exist { // 已经存在
			continue
		}
		newTableIdList = append(newTableIdList, params.TableIdList[i])
	}

	newDleTableIdList := make([]string, 0)
	for i := 0; i < len(rows); i++ {
		exist := false
		for j := 0; j < len(newTableIdList); j++ {
			exist = rows[i].TableId == newTableIdList[j]
			if exist {
				break
			}
		}
		if exist {
			continue
		}
		newDleTableIdList = append(newDleTableIdList, rows[i].TableId)
	}

	spider.CorrelationBatchAdditional(params.TableType, params.RoleId, newTableIdList, newDleTableIdList)
	service.Success()

}
