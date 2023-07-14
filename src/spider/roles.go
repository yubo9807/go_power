package spider

import (
	"server/configs"
	"server/src/service"
	"server/src/utils"
	"strings"
	"time"
)

var Roles rolesTable

type rolesTable struct {
	dbKeys  []string
	keysOwn string
}

func init() {
	Roles.dbKeys = utils.GetStructDBKeys(RoleColumn{})
	Roles.keysOwn = strings.Join(Roles.dbKeys, ", ")
}

type RoleColumn struct {
	Id         string  `json:"id"`
	Role       string  `json:"role"`
	CreateTime int     `json:"createTime" db:"create_time"`
	UpdateTime *int    `json:"updateTime" db:"update_time"`
	Remark     *string `json:"remark"`
}

// 获取角色列表
func (r *rolesTable) RoleList(role string) []RoleColumn {
	db := service.Sql.DBConnect()
	defer db.Close()
	likeStr := utils.If(role == "", "", " WHERE role = '"+role+"'")
	var roleList []RoleColumn
	sqlStr := "SELECT " + r.keysOwn + " FROM " + configs.Table_Roles + likeStr + " ORDER BY create_time ASC;"
	err := db.Select(&roleList, sqlStr)
	if err != nil {
		panic(err.Error())
	}
	return roleList
}

// 添加角色
func (r *rolesTable) Additional(role, remark string) {
	db := service.Sql.DBConnect()
	defer db.Close()
	id := utils.CreateID()
	createTime := time.Now().Unix()
	_, err := db.Exec("INSERT INTO "+configs.Table_Roles+"(id, role, remark, create_time) values(?, ?, ?, ?);",
		id, role, remark, createTime)
	if err != nil {
		panic(err.Error())
	}
}

// 修改信息
func (r *rolesTable) Update(id, role, remark string) {
	db := service.Sql.DBConnect()
	defer db.Close()
	updateTime := time.Now().Unix()
	_, err := db.Exec("UPDATE "+configs.Table_Roles+" SET update_time = ?, role = ?, remark = ? WHERE id = ?;",
		updateTime, role, remark, id)
	if err != nil {
		panic(err.Error())
	}
}
