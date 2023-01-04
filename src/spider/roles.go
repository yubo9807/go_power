package spider

import (
	"server/src/service"
	"server/src/utils"
	"time"
)

type Role struct {
	Id         string  `json:"id"`
	Role       string  `json:"role"`
	CreateTime int     `json:"createTime" db:"create_time"`
	UpdateTime *int    `json:"updateTime" db:"update_time"`
	Remark     *string `json:"remark"`
}

// 获取角色列表
func RoleList() []Role {
	db := service.DBConnect()
	defer db.Close()
	var roleList []Role
	err := db.Select(&roleList, "SELECT * FROM roles;")
	if err != nil {
		panic(err.Error())
	}
	return roleList
}

func RoleQuery(role string) []Role {
	db := service.DBConnect()
	defer db.Close()
	var roleList []Role
	err := db.Select(&roleList, "SELECT * FROM element WHERE name = '"+role+"';")
	if err != nil {
		panic(err.Error())
	}
	return roleList
}

// 添加角色
func RoleAdditional(role string) {
	db := service.DBConnect()
	defer db.Close()
	id := utils.CreateID()
	createTime := time.Now().Unix()
	_, err := db.Exec("INSERT INTO roles(id, role, create_time) values(?, ?, ?);", id, role, createTime)
	if err != nil {
		panic(err.Error())
	}
}

// 修改信息
func RoleUpdate(id, role, remark string) {
	db := service.DBConnect()
	defer db.Close()
	updateTime := time.Now().Unix()
	_, err := db.Exec(`UPDATE roles SET update_time = ?, role = ?, remark = ? WHERE id = ?;`, updateTime, role, remark, id)
	if err != nil {
		panic(err.Error())
	}
}
