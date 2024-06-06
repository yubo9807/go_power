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
	Parent     *string `json:"parent"`
	Remark     *string `json:"remark"`
}

// 获取角色列表
func (r *rolesTable) RoleList(role string) []RoleColumn {
	db := service.Sql.DBConnect()
	defer db.Close()
	var roleList []RoleColumn
	var sqlStr string
	if role == "" {
		sqlStr = "SELECT " + r.keysOwn + " FROM " + configs.Table_Roles + " ORDER BY create_time ASC;"
		err := db.Select(&roleList, sqlStr)
		if err != nil {
			panic(err.Error())
		}
	} else {
		sqlStr = "SELECT " + r.keysOwn + " FROM " + configs.Table_Roles + " WHERE role = ? ORDER BY create_time ASC;"
		err := db.Select(&roleList, sqlStr, role)
		if err != nil {
			panic(err.Error())
		}
	}
	return roleList
}

// 添加角色
func (r *rolesTable) Additional(role string, remark, parent *string) {
	db := service.Sql.DBConnect()
	defer db.Close()
	id := utils.CreateID()
	createTime := time.Now().Unix()
	_, err := db.Exec("INSERT INTO "+configs.Table_Roles+"(id, role, remark, parent, create_time) values(?, ?, ?, ?, ?);",
		id, role, remark, parent, createTime)
	if err != nil {
		panic(err.Error())
	}
}

// 修改信息
func (r *rolesTable) Update(id, role string, remark, parent *string) {
	db := service.Sql.DBConnect()
	defer db.Close()
	updateTime := time.Now().Unix()
	_, err := db.Exec("UPDATE "+configs.Table_Roles+" SET update_time = ?, role = ?, remark = ?, parent = ? WHERE id = ?;",
		updateTime, role, remark, parent, id)
	if err != nil {
		panic(err.Error())
	}
}

func (r *rolesTable) Delete(id string) {
	db := service.Sql.DBConnect()
	defer db.Close()

	// 删除自己
	{
		_, err := db.Exec("DELETE FROM "+configs.Table_Roles+" WHERE id = ?;", id)
		if err != nil {
			panic(err.Error())
		}
	}

	// 关联表相关数据
	{
		_, err := db.Exec("DELETE FROM "+configs.Table_Correlation+" WHERE role_id = ?;", id)
		if err != nil {
			panic(err.Error())
		}
	}

	// 找子级
	{
		var roleList []RoleColumn
		sqlStr := "SELECT id FROM " + configs.Table_Roles + " WHERE parent = ?;"
		err := db.Select(&roleList, sqlStr, id)
		if err != nil {
			panic(err.Error())
		}
		for _, val := range roleList {
			r.Delete(val.Id)
		}
	}
}
