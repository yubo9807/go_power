package spider

import (
	"server/configs"
	"server/src/service"
	"server/src/utils"
	"strings"
	"time"
)

type interfaceTable struct {
	dbKeys  []string
	keysOwn string
}

var Interface interfaceTable

func init() {
	Interface.dbKeys = utils.GetStructDBKeys(tableInterfaceColumn{})
	Interface.keysOwn = strings.Join(Interface.dbKeys, ", ")
}

type tableInterfaceColumn struct {
	Id         string  `json:"id"`
	Method     string  `json:"method"`
	Url        string  `json:"url"`
	Name       string  `json:"name"`
	CreateTime int     `json:"createTime" db:"create_time"`
	UpdateTime *int    `json:"updateTime" db:"update_time"`
	MenuId     *string `json:"menuId" db:"menu_id"`
}
type InterfaceColumn struct {
	tableInterfaceColumn
	CorrelationId *string `json:"correlationId" db:"correlation_id"`
	RoleId        *string `json:"roleId" db:"role_id"`
	Selected      bool    `json:"selected"`
}

// 获取接口（按模块）
func (i *interfaceTable) List(menuId, url string) []InterfaceColumn {
	db := service.Sql.DBConnect()
	defer db.Close()
	var interfaceList []InterfaceColumn
	newUrl := "%" + url + "%"
	if menuId == "" {
		err := db.Select(&interfaceList, "SELECT "+i.keysOwn+" FROM "+configs.Table_Interface+" WHERE menu_id IS NULL AND url LIKE ?;", newUrl)
		if err != nil {
			panic(err.Error())
		}
	} else {
		err := db.Select(&interfaceList, "SELECT "+i.keysOwn+" FROM "+configs.Table_Interface+" WHERE menu_id = ? AND url LIKE ?;", menuId, newUrl)
		if err != nil {
			panic(err.Error())
		}
	}
	return interfaceList
}

// 获取有权限的接口（按模块）
// @param point == "" 时查询公共模块接口
func (i *interfaceTable) PowerListModule(roleId, menuId string) []InterfaceColumn {
	db := service.Sql.DBConnect()
	defer db.Close()
	var interfaceList []InterfaceColumn
	sqlStr := `SELECT
	t1.id, t1.method, t1.url,
	t2.id AS 'correlation_id',
	t3.id AS 'role_id'
	FROM ` + configs.Table_Interface + ` AS t1
	LEFT JOIN ` + configs.Table_Correlation + ` AS t2
	ON t1.id = t2.table_id
	LEFT JOIN ` + configs.Table_Roles + ` AS t3
	ON t2.role_id = t3.id
	WHERE t2.table_type = 'interface' AND t3.id = ? AND t1.menu_id `
	if menuId == "" {
		err := db.Select(&interfaceList, sqlStr+"IS NULL;", roleId)
		if err != nil {
			panic(err.Error())
		}
	} else {
		err := db.Select(&interfaceList, sqlStr+"= ?;", roleId, menuId)
		if err != nil {
			panic(err.Error())
		}
	}
	return interfaceList
}

// 查询接口
func (i *interfaceTable) Query(method, url string) []InterfaceColumn {
	db := service.Sql.DBConnect()
	defer db.Close()
	var interfaceList []InterfaceColumn
	err := db.Select(&interfaceList, "SELECT "+i.keysOwn+" FROM "+configs.Table_Interface+" WHERE method = '"+method+"' AND url LIKE '%"+url+"%';")
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
	createTime := time.Now().Unix()
	_, err := db.Exec("INSERT INTO "+configs.Table_Interface+"(id, method, url, name, menu_id, create_time) values(?, ?, ?, ?, ?, ?);",
		id, method, url, name, menuId, createTime)
	if err != nil {
		panic(err.Error())
	}
}

// 获取具有权限的接口
func (i *interfaceTable) PowerList(roleId, method, url string) []InterfaceColumn {
	db := service.Sql.DBConnect()
	defer db.Close()
	dbKeys := utils.Map(i.dbKeys, func(val string, i int) string {
		return "t1." + val
	})
	keysOwn := strings.Join(dbKeys, ", ")
	var interfaceList []InterfaceColumn
	err := db.Select(&interfaceList, `SELECT `+keysOwn+` FROM `+configs.Table_Interface+` AS t1
	LEFT JOIN `+configs.Table_Correlation+` AS t2 ON t1.id = t2.table_id
	LEFT JOIN `+configs.Table_Roles+` AS t3 ON t2.role_id = t3.id
	WHERE t2.table_type = 'interface' AND t3.id = ? AND method = ? AND url = ?;`,
		roleId, method, url)
	if err != nil {
		panic(err.Error())
	}
	return interfaceList
}
