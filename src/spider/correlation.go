package spider

import (
	"server/src/service"
	"server/src/utils"
	"time"
)

type Correlation struct {
	Id         string `json:"id"`
	RoleId     string `json:"roleId" db:"role_id"`
	TableId    string `json:"tableId" db:"table_id"`
	TableType  string `json:"tableType" db:"table_type"`
	CreateTime int    `json:"createTime" db:"create_time"`
	UpdateTime *int   `json:"updateTime" db:"update_time"`
}

// 追加关联关系，给指定的角色添加权限
func CorrelationAdditional(roleId, tableId, tableType string) {
	db := service.DBConnect()
	defer db.Close()
	id := utils.CreateID()
	createTime := time.Now().Unix()
	_, err := db.Exec("INSERT INTO correlation(id, role_id, table_id, table_type, create_time) values(?, ?, ?, ?, ?);",
		id, roleId, tableId, tableType, createTime)
	if err != nil {
		panic(err.Error())
	}
}

// 批量同步关联关系
func CorrelationBatchAdditional(tableType, roleId string, tableIdList, delTableIdList []string) {
	db := service.DBConnect()
	defer db.Close()

	// 添加
	for i := 0; i < len(tableIdList); i++ {
		id := utils.CreateID()
		createTime := time.Now().Unix()
		_, err := db.Exec("INSERT INTO correlation(id, role_id, table_id, table_type, create_time) values(?, ?, ?, ?, ?);",
			id, roleId, tableIdList[i], tableType, createTime)
		if err != nil {
			panic(err.Error())
		}
	}

	// 删除
	for i := 0; i < len(delTableIdList); i++ {
		_, err := db.Exec("DELETE FROM correlation WHERE table_id = ? AND role_id = ?;", delTableIdList[i], roleId)
		if err != nil {
			panic(err.Error())
		}
	}
}

// 删除关联的数据
func CorrelationDeleteCorrelation(tableType, tableId string) {
	db := service.DBConnect()
	defer db.Close()
	_, err := db.Exec("DELETE FROM correlation WHERE table_type = ? AND table_id = ?;", tableType, tableId)
	if err != nil {
		panic(err.Error())
	}
}

// 按类型查询
func CorrelationTableTypeQuery(roleId, tableType string) []Correlation {
	db := service.DBConnect()
	defer db.Close()
	var correlation []Correlation
	err := db.Select(&correlation, "SELECT * FROM correlation WHERE role_id = '"+roleId+"' AND table_type = '"+tableType+"';")
	if err != nil {
		panic(err.Error())
	}
	return correlation
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
