package spider

import (
	"server/configs"
	"server/src/service"
	"server/src/utils"
	"strings"
	"time"
)

type elementTable struct {
	dbKeys  []string
	keysOwn string
}

var Elememt elementTable

func init() {
	Elememt.dbKeys = utils.GetStructDBKeys(tableElementColumn{})
	newKeys := utils.Map(Elememt.dbKeys, func(val string, i int) string {
		if val == "key" {
			val = "`" + val + "`"
		}
		return val
	})
	Elememt.keysOwn = strings.Join(newKeys, ", ")
}

type tableElementColumn struct {
	Id         string  `json:"id"`
	Key        string  `json:"key"`
	Name       string  `json:"name"`
	CreateTime int     `json:"createTime" db:"create_time"`
	UpdateTime *int    `json:"updateTime" db:"update_time"`
	MenuId     *string `json:"menuId" db:"menu_id"`
}
type ElememtColumn struct {
	tableElementColumn
	CorrelationId *string `json:"correlationId" db:"correlation_id"`
	RoleId        *string `json:"roleId" db:"role_id"`
	Selected      bool    `json:"selected"`
}

// 获取所有元素
func (e *elementTable) List(menuId string) []ElememtColumn {
	db := service.Sql.DBConnect()
	defer db.Close()
	var elementList []ElememtColumn
	joint := "IS NULL"
	if menuId != "" {
		joint = "= '" + menuId + "'"
	}
	err := db.Select(&elementList, "SELECT "+e.keysOwn+" FROM "+configs.Table_Element+" WHERE menu_id "+joint+";")
	if err != nil {
		panic(err.Error())
	}
	return elementList
}

// 获取有权限的元素
// @param point == "" 时查询公共模块元素
func (e *elementTable) PowerList(roleId, menuId string) []ElememtColumn {
	db := service.Sql.DBConnect()
	defer db.Close()
	var elementList []ElememtColumn
	joint := "IS NULL"
	if menuId != "" {
		joint = "= '" + menuId + "'"
	}
	err := db.Select(&elementList, `SELECT
	t1.id, t1.key,
	t2.id AS 'correlation_id',
	t3.id AS 'role_id'
	FROM `+configs.Table_Element+` AS t1
	LEFT JOIN `+configs.Table_Correlation+` AS t2
	ON t1.id = t2.table_id
	LEFT JOIN `+configs.Table_Roles+` AS t3
	ON t2.role_id = t3.id
	WHERE t2.table_type = 'element' AND t3.id = '`+roleId+"' AND t1.menu_id "+joint+";")
	if err != nil {
		panic(err.Error())
	}
	return elementList
}

// 获取有权限的元素
func (e *elementTable) PowerList2(roleId string) []ElememtColumn {
	db := service.Sql.DBConnect()
	defer db.Close()
	var elementList []ElememtColumn
	err := db.Select(&elementList, `SELECT
	t1.id, t1.key,
	t2.id AS 'correlation_id',
	t3.id AS 'role_id'
	FROM `+configs.Table_Element+` AS t1
	LEFT JOIN `+configs.Table_Correlation+` AS t2
	ON t1.id = t2.table_id
	LEFT JOIN `+configs.Table_Roles+` AS t3
	ON t2.role_id = t3.id
	WHERE t2.table_type = 'element' AND t3.id = '`+roleId+"';")
	if err != nil {
		panic(err.Error())
	}
	return elementList
}

// 查询元素
func (e *elementTable) Query(key, name string) []ElememtColumn {
	db := service.Sql.DBConnect()
	defer db.Close()
	var elementList []ElememtColumn
	err := db.Select(&elementList, "SELECT "+e.keysOwn+" FROM "+configs.Table_Element+" WHERE 'key' = '"+key+"' AND 'name' LIKE '%"+name+"%';")
	if err != nil {
		panic(err.Error())
	}
	return elementList
}

// 修改元素数据
func (e *elementTable) Modify(id, key, name, menuId string) {
	db := service.Sql.DBConnect()
	defer db.Close()
	updateTime := time.Now().Unix()
	_, err := db.Exec("UPDATE "+configs.Table_Element+" SET update_time = ?, key = ?, name = ?, menu_id = ? WHERE id = ?;",
		updateTime, key, name, menuId, id)
	if err != nil {
		panic(err.Error())
	}
}

// 添加元素
func (e *elementTable) Additional(key, name string, menuId *string) {
	db := service.Sql.DBConnect()
	defer db.Close()
	id := utils.CreateID()
	createTime := time.Now().Unix()
	_, err := db.Exec("INSERT INTO "+configs.Table_Element+"(id, `key`, `name`, menu_id, create_time) values(?, ?, ?, ?, ?);",
		id, key, name, menuId, createTime)
	if err != nil {
		panic(err.Error())
	}
}
