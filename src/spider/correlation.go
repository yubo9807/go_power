package spider

import (
	"server/src/service"
	"server/src/utils"
	"time"
)

type Correlation struct {
	Id         string `json:"id"`
	Role_id    string `json:"role_id"`
	Table_id   string `json:"table_id"`
	Table_type string `json:"table_type"`
	CreateTime int    `json:"createTime" db:"create_time"`
	UpdateTime *int   `json:"updateTime" db:"update_time"`
}

// 追加关联关系，给指定的角色添加权限
func CorrelationAdditional(roleId, tableId, tableType string) {
	db := service.DBConnect()
	defer db.Close()
	id := utils.CreateID()
	createTime := time.Now().Unix()
	_, err := db.Exec("INSERT INTO correlation(id, role_id, table_id, table_type, create_time) values(?, ?, ?);", id, roleId, tableId, tableType, createTime)
	if err != nil {
		panic(err.Error())
	}
}

// 删除符合关联 table_id 的数据
func CorrelationDeleteCorrelation(tableId string) {
	db := service.DBConnect()
	defer db.Close()
	_, err := db.Exec("DELETE FROM correlation WHERE table_id = ?;", tableId)
	if err != nil {
		panic(err.Error())
	}
}

// 查询已存在的关联
func CorrelationQuery(roleId, tableId, tableType string) []Correlation {
	db := service.DBConnect()
	defer db.Close()
	var correlation []Correlation
	err := db.Select(&correlation, "SELECT * FROM correlation WHERE role_id = "+roleId+" AND table_id = "+tableId+" AND table_type = '"+tableType+"';")
	if err != nil {
		panic(err.Error())
	}
	return correlation
}
