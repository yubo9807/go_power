package spider

import (
	"server/src/service"
	"server/src/utils"
	"time"
)

type Interface struct {
	Id         string  `json:"id"`
	Url        string  `json:"url"`
	Name       string  `json:"name"`
	CreateTime int     `json:"createTime" db:"create_time"`
	UpdateTime *int    `json:"updateTime" db:"update_time"`
	Point      *string `json:"point"`
	Selected   bool    `json:"selected"`
}

// 获取所有接口
func InterfaceList(point string) []Interface {
	db := service.DBConnect()
	defer db.Close()
	var interfaceList []Interface
	joint := "IS NULL"
	if point != "" {
		joint = "= '" + point + "'"
	}
	err := db.Select(&interfaceList, "SELECT * FROM interface WHERE point "+joint+";")
	if err != nil {
		panic(err.Error())
	}
	return interfaceList
}

// 获取有权限的接口
// @param point == "" 时查询公共模块接口
func InterfacePowerList(role, point string) []Interface {
	db := service.DBConnect()
	defer db.Close()
	var interfaceList []Interface
	joint := "IS NULL"
	if point != "" {
		joint = "= '" + point + "'"
	}
	err := db.Select(&interfaceList, `SELECT t1.* FROM interface AS t1
	LEFT JOIN correlation AS t2
	ON t1.id = t2.table_id
	LEFT JOIN roles AS t3
	ON t2.role_id = t3.id
	WHERE t2.table_name = 'interface' AND t3.role = '`+role+"' AND t1.point "+joint+";")
	if err != nil {
		panic(err.Error())
	}
	return interfaceList
}

// 查询接口
func InterfaceQuery(url string) []Interface {
	db := service.DBConnect()
	defer db.Close()
	var interfaceList []Interface
	err := db.Select(&interfaceList, "SELECT * FROM interface WHERE name LIKE '%"+url+"%';")
	if err != nil {
		panic(err.Error())
	}
	return interfaceList
}

// 修改接口数据
func InterfaceModify(id, url string) {
	db := service.DBConnect()
	defer db.Close()
	updateTime := time.Now().Unix()
	_, err := db.Exec(`UPDATE interface SET update_time = ?, url = ? WHERE id = ?;`, updateTime, url, id)
	if err != nil {
		panic(err.Error())
	}
}

func InterfaceModifyMenu(id, menu string) {
	db := service.DBConnect()
	defer db.Close()
	updateTime := time.Now().Unix()
	_, err := db.Exec(`UPDATE interface SET update_time = ?, menu = ? WHERE id = ?;`, updateTime, menu, id)
	if err != nil {
		panic(err.Error())
	}
}

// 添加接口
func InterfaceAdditional(url, name string) {
	db := service.DBConnect()
	defer db.Close()
	id := utils.CreateID()
	createTime := time.Now().Unix()
	_, err := db.Exec("INSERT INTO interface(id, url, name create_time) values(?, ?, ?);", id, url, name, createTime)
	if err != nil {
		panic(err.Error())
	}
}
