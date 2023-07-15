package spider

import (
	"server/configs"
	"server/src/service"
	"server/src/utils"
	"strings"
	"time"
)

type correlationTable struct {
	dbKeys  []string
	keysOwn string
}

var Correlation correlationTable

func init() {
	Correlation.dbKeys = utils.GetStructDBKeys(CorrelationColumn{})
	Correlation.keysOwn = strings.Join(Correlation.dbKeys, ", ")
}

type CorrelationColumn struct {
	Id         string `json:"id"`
	RoleId     string `json:"roleId" db:"role_id"`
	TableId    string `json:"tableId" db:"table_id"`
	TableType  string `json:"tableType" db:"table_type"`
	CreateTime int    `json:"createTime" db:"create_time"`
	UpdateTime *int   `json:"updateTime" db:"update_time"`
}

// 追加关联关系，给指定的角色添加权限
func (c *correlationTable) Additional(roleId, tableId, tableType string) {
	db := service.Sql.DBConnect()
	defer db.Close()
	id := utils.CreateID()
	createTime := time.Now().Unix()
	_, err := db.Exec("INSERT INTO "+configs.Table_Correlation+"(id, role_id, table_id, table_type, create_time) values(?, ?, ?, ?, ?);",
		id, roleId, tableId, tableType, createTime)
	if err != nil {
		panic(err.Error())
	}
}

// 批量同步关联关系
func (c *correlationTable) BatchAdditional(tableType, roleId string, tableIdList, delTableIdList []string) {
	db := service.Sql.DBConnect()
	defer db.Close()

	// 添加
	for i := 0; i < len(tableIdList); i++ {
		id := utils.CreateID()
		createTime := time.Now().Unix()
		_, err := db.Exec("INSERT INTO "+configs.Table_Correlation+"(id, role_id, table_id, table_type, create_time) values(?, ?, ?, ?, ?);",
			id, roleId, tableIdList[i], tableType, createTime)
		if err != nil {
			panic(err.Error())
		}
	}

	// 删除
	for i := 0; i < len(delTableIdList); i++ {
		_, err := db.Exec("DELETE FROM "+configs.Table_Correlation+" WHERE table_id = ? AND role_id = ?;",
			delTableIdList[i], roleId)
		if err != nil {
			panic(err.Error())
		}
	}
}

// 删除关联的数据
func (c *correlationTable) DeleteCorrelation(tableType, tableId string) {
	db := service.Sql.DBConnect()
	defer db.Close()
	_, err := db.Exec("DELETE FROM "+configs.Table_Correlation+" WHERE table_type = ? AND table_id = ?;",
		tableType, tableId)
	if err != nil {
		panic(err.Error())
	}
}

// 按类型查询
func (c *correlationTable) TableTypeQuery(roleId, tableType string) []CorrelationColumn {
	db := service.Sql.DBConnect()
	defer db.Close()
	var correlation []CorrelationColumn
	err := db.Select(&correlation, "SELECT "+c.keysOwn+" FROM "+configs.Table_Correlation+" WHERE role_id = ? AND table_type = ?;",
		roleId, tableType)
	if err != nil {
		panic(err.Error())
	}
	return correlation
}

// 查询已存在的关联
func (c *correlationTable) Query(roleId, tableId, tableType string) []CorrelationColumn {
	db := service.Sql.DBConnect()
	defer db.Close()
	var correlation []CorrelationColumn
	err := db.Select(&correlation, "SELECT "+c.keysOwn+" FROM "+configs.Table_Correlation+" WHERE role_id = ? AND table_id = ? AND table_type = ?;",
		roleId, tableId, tableType)
	if err != nil {
		panic(err.Error())
	}
	return correlation
}
