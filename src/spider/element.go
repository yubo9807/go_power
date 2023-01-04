package spider

import (
	"server/src/service"
	"server/src/utils"
	"time"
)

type Elememt struct {
	Id         string  `json:"id"`
	Key        string  `json:"key"`
	Name       string  `json:"name"`
	CreateTime int     `json:"createTime" db:"create_time"`
	UpdateTime *int    `json:"updateTime" db:"update_time"`
	MenuId     *string `json:"menuId" db:"menu_id"`

	CorrelationId *string `json:"correlationId" db:"correlation_id"`
	RoleId        *string `json:"roleId" db:"role_id"`
	Selected      bool    `json:"selected"`
}

// 获取所有接口
func ElememtList(menuId string) []Elememt {
	db := service.DBConnect()
	defer db.Close()
	var elementList []Elememt
	joint := "IS NULL"
	if menuId != "" {
		joint = "= '" + menuId + "'"
	}
	err := db.Select(&elementList, "SELECT * FROM element WHERE menu_id "+joint+";")
	if err != nil {
		panic(err.Error())
	}
	return elementList
}

// 获取有权限的元素
// @param point == "" 时查询公共模块元素
func ElememtPowerList(roleId, menuId string) []Elememt {
	db := service.DBConnect()
	defer db.Close()
	var elementList []Elememt
	joint := "IS NULL"
	if menuId != "" {
		joint = "= '" + menuId + "'"
	}
	err := db.Select(&elementList, `SELECT
	t1.*,
	t2.id AS 'correlation_id',
	t3.id AS 'role_id'
	FROM element AS t1
	LEFT JOIN correlation AS t2
	ON t1.id = t2.table_id
	LEFT JOIN roles AS t3
	ON t2.role_id = t3.id
	WHERE t2.table_type = 'element' AND t3.id = '`+roleId+"' AND t1.menu_id "+joint+";")
	if err != nil {
		panic(err.Error())
	}
	return elementList
}

// 查询元素
func ElememtQuery(key, name string) []Elememt {
	db := service.DBConnect()
	defer db.Close()
	var elementList []Elememt
	err := db.Select(&elementList, "SELECT * FROM element WHERE key = '"+key+"' AND name LIKE '%"+name+"%';")
	if err != nil {
		panic(err.Error())
	}
	return elementList
}

// 修改元素数据
func ElememtModify(id, key, name string) {
	db := service.DBConnect()
	defer db.Close()
	updateTime := time.Now().Unix()
	_, err := db.Exec(`UPDATE element SET update_time = ?, key = ? name = ? WHERE id = ?;`, updateTime, key, name, id)
	if err != nil {
		panic(err.Error())
	}
}

// 添加元素
func ElememtAdditional(key, name string) {
	db := service.DBConnect()
	defer db.Close()
	id := utils.CreateID()
	createTime := time.Now().Unix()
	_, err := db.Exec("INSERT INTO element(id, key, name, create_time) values(?, ?, ?, ?);", id, key, name, createTime)
	if err != nil {
		panic(err.Error())
	}
}
