package spider

import (
	"server/configs"
	"server/src/service"
	"server/src/utils"
	"time"
)

type rolesTable struct{}

var Roles rolesTable

type RoleColumn struct {
	Id         string  `json:"id"`
	Role       string  `json:"role"`
	CreateTime int     `json:"createTime" db:"create_time"`
	UpdateTime *int    `json:"updateTime" db:"update_time"`
	Remark     *string `json:"remark"`
}

// 获取角色列表
func (r *rolesTable) RoleList() []RoleColumn {
	db := service.Sql.DBConnect()
	defer db.Close()
	var roleList []RoleColumn
	err := db.Select(&roleList, "SELECT * FROM "+configs.Table_Roles+";")
	if err != nil {
		panic(err.Error())
	}
	return roleList
}

// 添加角色
func (r *rolesTable) Additional(role string) {
	db := service.Sql.DBConnect()
	defer db.Close()
	id := utils.CreateID()
	createTime := time.Now().Unix()
	_, err := db.Exec("INSERT INTO "+configs.Table_Roles+"(id, role, create_time) values(?, ?, ?);", id, role, createTime)
	if err != nil {
		panic(err.Error())
	}
}

// 修改信息
func (r *rolesTable) Update(id, role, remark string) {
	db := service.Sql.DBConnect()
	defer db.Close()
	updateTime := time.Now().Unix()
	_, err := db.Exec("UPDATE "+configs.Table_Roles+" SET update_time = ?, role = ?, remark = ? WHERE id = ?;", updateTime, role, remark, id)
	if err != nil {
		panic(err.Error())
	}
}
