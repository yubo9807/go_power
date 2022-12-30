package spider

import (
	"server/src/service"
	"server/src/utils"
	"time"
)

type Menu struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	CreateTime int     `json:"createTime" db:"create_time"`
	UpdateTime *int    `json:"updateTime" db:"update_time"`
	Title      *string `json:"title"`
	Hidden     bool    `json:"hidden"`
	Parent     *string `json:"parent"`
	Selected   bool    `json:"selected"`
}

// 获取所有菜单
func MenuList() []Menu {
	db := service.DBConnect()
	defer db.Close()
	var menuList []Menu
	err := db.Select(&menuList, "SELECT * FROM menu;")
	if err != nil {
		panic(err.Error())
	}
	return menuList
}

// 获取有权限的菜单
func MenuPowerList(role string) []Menu {
	db := service.DBConnect()
	defer db.Close()
	var menuList []Menu
	err := db.Select(&menuList, `SELECT t1.* FROM menu AS t1
	LEFT JOIN correlation AS t2
	ON t1.name = t2.name
	WHERE t2.type = 'menu' AND t2.role = '`+role+"';")
	if err != nil {
		panic(err.Error())
	}
	return menuList
}

// 查询菜单
func MenuQuery(name, title string) []Menu {
	db := service.DBConnect()
	defer db.Close()
	var menuList []Menu
	err := db.Select(&menuList, "SELECT * FROM menu WHERE name LIKE '%"+name+"%' AND title LIKE '%"+title+"%';")
	if err != nil {
		panic(err.Error())
	}
	return menuList
}

// 添加菜单
func MenuAdditional(name string) {
	db := service.DBConnect()
	defer db.Close()
	id := utils.CreateID()
	createTime := time.Now().Unix()
	_, err := db.Exec("INSERT INTO menu(id, name, create_time) values(?, ?, ?);", id, name, createTime)
	if err != nil {
		panic(err.Error())
	}
}

// 修改菜单数据
func MenuModify(id, name string, hidden bool, title string) {
	db := service.DBConnect()
	defer db.Close()
	updateTime := time.Now().Unix()
	_, err := db.Exec(`UPDATE menu SET update_time = ?, name = ?, hidden = ?, title = ? WHERE id = ?;`, updateTime, name, hidden, title, id)
	if err != nil {
		panic(err.Error())
	}
}
