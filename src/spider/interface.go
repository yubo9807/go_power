package spider

import (
	"fmt"
	"server/configs"
	"server/src/service"
	"server/src/utils"
	"time"
)

type interfaceTable struct{}

var Interface interfaceTable

type InterfaceColumn struct {
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

// 获取接口（按模块）
func (i *interfaceTable) List(menuId, url string) []InterfaceColumn {
	db := service.Sql.DBConnect()
	defer db.Close()
	var interfaceList []InterfaceColumn
	joint := "IS NULL"
	if menuId != "" {
		joint = "= '" + menuId + "'"
	}
	err := db.Select(&interfaceList, "SELECT * FROM "+configs.Table_Interface+" WHERE menu_id "+joint+" AND url LIKE '%"+url+"%';")
	if err != nil {
		panic(err.Error())
	}
	return interfaceList
}

// 获取有权限的接口（按模块）
// @param point == "" 时查询公共模块接口
func (i *interfaceTable) PowerListModule(roleId, menuId string) []InterfaceColumn {
	db := service.Sql.DBConnect()
	defer db.Close()
	var interfaceList []InterfaceColumn
	joint := "IS NULL"
	if menuId != "" {
		joint = "= '" + menuId + "'"
	}
	err := db.Select(&interfaceList, `SELECT
	t1.*,
	t2.id AS 'correlation_id',
	t3.id AS 'role_id'
	FROM `+configs.Table_Interface+` AS t1
	LEFT JOIN `+configs.Table_Correlation+` AS t2
	ON t1.id = t2.table_id
	LEFT JOIN `+configs.Table_Roles+` AS t3
	ON t2.role_id = t3.id
	WHERE t2.table_type = 'interface' AND t3.id = '`+roleId+"' AND t1.menu_id "+joint+";")
	if err != nil {
		panic(err.Error())
	}
	return interfaceList
}

// 查询接口
func (i *interfaceTable) Query(method, url string) []InterfaceColumn {
	db := service.Sql.DBConnect()
	defer db.Close()
	var interfaceList []InterfaceColumn
	err := db.Select(&interfaceList, "SELECT * FROM "+configs.Table_Interface+" WHERE method = '"+method+"' AND url LIKE '%"+url+"%';")
	if err != nil {
		panic(err.Error())
	}
	return interfaceList
}

// 修改接口数据
func (i *interfaceTable) Modify(id, method, url, name string, menuId *string) {
	db := service.Sql.DBConnect()
	defer db.Close()
	updateTime := time.Now().Unix()
	_, err := db.Exec("UPDATE "+configs.Table_Interface+" SET update_time = ?, method = ?, url = ?, name = ?, menu_id = ? WHERE id = ?;",
		updateTime, method, url, name, menuId, id)
	if err != nil {
		panic(err.Error())
	}
}

// 添加接口
func (i *interfaceTable) Additional(method, url, name string, menuId *string) {
	db := service.Sql.DBConnect()
	defer db.Close()
	id := utils.CreateID()
	fmt.Println(id)
	createTime := time.Now().Unix()
	_, err := db.Exec("INSERT INTO "+configs.Table_Interface+"(id, method, url, name, menu_id, create_time) values(?, ?, ?, ?, ?, ?);",
		id, method, url, name, menuId, createTime)
	if err != nil {
		panic(err.Error())
	}
}

// 获取具有权限的接口
// method, url 不为空时进行精确查询
func (i *interfaceTable) PowerList(roleId, method, url string) []InterfaceColumn {
	joint := ";"
	if method != "" && url != "" {
		joint = " AND method = '" + method + "' AND url = '" + url + "';"
	}
	db := service.Sql.DBConnect()
	defer db.Close()
	var interfaceList []InterfaceColumn
	err := db.Select(&interfaceList, `SELECT t1.* FROM `+configs.Table_Interface+` AS t1
	LEFT JOIN `+configs.Table_Correlation+` AS t2 ON t1.id = t2.table_id
	LEFT JOIN `+configs.Table_Roles+` AS t3 ON t2.role_id = t3.id
	WHERE t2.table_type = 'interface' AND t3.id = '`+roleId+`'`+joint)
	if err != nil {
		panic(err.Error())
	}
	return interfaceList
}
