package spider

import (
	"fmt"
	"server/src/service"
	"server/src/utils"
	"time"
)

type Interface struct {
	Id         string  `json:"id"`
	Method     string  `json:"method"`
	Url        string  `json:"url"`
	Name       string  `json:"name"`
	CreateTime int     `json:"createTime" db:"create_time"`
	UpdateTime *int    `json:"updateTime" db:"update_time"`
	MenuId     *string `json:"menuId" db:"menu_id"`

	CorrelationId *string `json:"correlationId" db:"correlation_id"`
	RoleId        *string `json:"roleId" db:"role_id"`
	Selected      bool    `json:"selected"`
}

// 获取所有接口
func InterfaceList(menuId string) []Interface {
	db := service.DBConnect()
	defer db.Close()
	var interfaceList []Interface
	joint := "IS NULL"
	if menuId != "" {
		joint = "= '" + menuId + "'"
	}
	err := db.Select(&interfaceList, "SELECT * FROM interface WHERE menu_id "+joint+";")
	if err != nil {
		panic(err.Error())
	}
	return interfaceList
}

// 获取有权限的接口
// @param point == "" 时查询公共模块接口
func InterfacePowerList(role, menuId string) []Interface {
	db := service.DBConnect()
	defer db.Close()
	var interfaceList []Interface
	joint := "IS NULL"
	if menuId != "" {
		joint = "= '" + menuId + "'"
	}
	err := db.Select(&interfaceList, `SELECT
	t1.*,
	t2.id AS 'correlation_id',
	t3.id AS 'role_id'
	FROM interface AS t1
	LEFT JOIN correlation AS t2
	ON t1.id = t2.table_id
	LEFT JOIN roles AS t3
	ON t2.role_id = t3.id
	WHERE t2.table_type = 'interface' AND t3.role = '`+role+"' AND t1.menu_id "+joint+";")
	if err != nil {
		panic(err.Error())
	}
	return interfaceList
}

// 查询接口
func InterfaceQuery(method, url string) []Interface {
	db := service.DBConnect()
	defer db.Close()
	var interfaceList []Interface
	err := db.Select(&interfaceList, "SELECT * FROM interface WHERE method = '"+method+"' AND url LIKE '%"+url+"%';")
	if err != nil {
		panic(err.Error())
	}
	return interfaceList
}

// 修改接口数据
func InterfaceModify(id, method, url, name string, menuId *string) {
	db := service.DBConnect()
	defer db.Close()
	updateTime := time.Now().Unix()
	_, err := db.Exec(`UPDATE interface SET update_time = ?, method = ?, url = ?, name = ?, menu_id = ? WHERE id = ?;`,
		updateTime, method, url, name, menuId, id)
	if err != nil {
		panic(err.Error())
	}
}

// 添加接口
func InterfaceAdditional(method, url, name string, menuId *string) {
	db := service.DBConnect()
	defer db.Close()
	id := utils.CreateID()
	fmt.Println(id)
	createTime := time.Now().Unix()
	_, err := db.Exec("INSERT INTO interface(id, method, url, name, menu_id, create_time) values(?, ?, ?, ?, ?, ?);",
		id, method, url, name, menuId, createTime)
	if err != nil {
		panic(err.Error())
	}
}
