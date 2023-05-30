package spider

import (
	"server/configs"
	"server/src/service"
	"server/src/utils"
	"time"
)

type menuTable struct{}

var Menu menuTable

type MenuColumn struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	CreateTime int     `json:"createTime" db:"create_time"`
	UpdateTime *int    `json:"updateTime" db:"update_time"`
	Title      *string `json:"title"`
	Hidden     bool    `json:"hidden"`
	Parent     *string `json:"parent"`
	Count      int     `json:"count"`

	CorrelationId *string `json:"correlationId" db:"correlation_id"`
	RoleId        *string `json:"roleId" db:"role_id"`
	Selected      bool    `json:"selected"`
}

// 获取所有菜单
func (m *menuTable) List(title string) []MenuColumn {
	db := service.Sql.DBConnect()
	defer db.Close()
	var menuList []MenuColumn
	err := db.Select(&menuList, "SELECT * FROM "+configs.Table_Menu+" WHERE title LIKE '%"+title+"%' ORDER BY count ASC;")
	if err != nil {
		panic(err.Error())
	}
	return menuList
}

// 获取有权限的菜单
func (m *menuTable) PowerList(roleId string) []MenuColumn {
	db := service.Sql.DBConnect()
	defer db.Close()
	var menuList []MenuColumn
	err := db.Select(&menuList, `SELECT
	t1.*, t2.id AS 'correlation_id', t3.id AS 'role_id'
	FROM `+configs.Table_Menu+` AS t1
	LEFT JOIN `+configs.Table_Correlation+` AS t2
	ON t1.id = t2.table_id
	LEFT JOIN `+configs.Table_Roles+` AS t3
	ON t2.role_id = t3.id
	WHERE t2.table_type = 'menu' AND t3.id = '`+roleId+"';")
	if err != nil {
		panic(err.Error())
	}
	return menuList
}

// 查询菜单
func (m *menuTable) Query(name, title string) []MenuColumn {
	db := service.Sql.DBConnect()
	defer db.Close()
	var menuList []MenuColumn
	err := db.Select(&menuList, "SELECT * FROM "+configs.Table_Menu+" WHERE name LIKE '%"+name+"%' AND title LIKE '%"+title+"%';")
	if err != nil {
		panic(err.Error())
	}
	return menuList
}

// 结构查询
func (m *menuTable) StructureQuery(parent *string) []MenuColumn {
	db := service.Sql.DBConnect()
	defer db.Close()
	var menuList []MenuColumn
	err := db.Select(&menuList, "SELECT * FROM "+configs.Table_Menu+" WHERE id = "+*parent+";")
	if err != nil {
		panic(err.Error())
	}
	return menuList
}

// 添加菜单
func (m *menuTable) Additional(name, title string, hidden bool, parent *string) {
	db := service.Sql.DBConnect()
	defer db.Close()
	maxCount := 0
	{
		// 查询最大排序
		type menuCount struct {
			Count int
		}
		var menuCountList []menuCount
		err := db.Select(&menuCountList, "SELECT MAX(`count`) AS count FROM "+configs.Table_Menu+";")
		if err != nil {
			panic(err.Error())
		}
		if len(menuCountList) > 0 {
			maxCount = menuCountList[0].Count + 1
		}
	}
	id := utils.CreateID()
	createTime := time.Now().Unix()
	_, err := db.Exec("INSERT INTO "+configs.Table_Menu+"(id, name, title, hidden, parent, count, create_time) values(?, ?, ?, ?, ?, ?, ?);",
		id, name, title, hidden, parent, maxCount, createTime)
	if err != nil {
		panic(err.Error())
	}
}

// 修改菜单数据
func (m *menuTable) Modify(id, name, title string, hidden bool, parent *string) {
	db := service.Sql.DBConnect()
	defer db.Close()
	updateTime := time.Now().Unix()
	_, err := db.Exec("UPDATE "+configs.Table_Menu+" SET update_time = ?, name = ?, title = ?, hidden = ?, parent = ? WHERE id = ?;", updateTime, name, title, hidden, parent, id)
	if err != nil {
		panic(err.Error())
	}
}
