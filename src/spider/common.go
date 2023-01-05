package spider

import (
	"server/src/service"
	"strconv"
	"time"
)

type commonType struct{}

var Common commonType

// 删除某张表中的某条数据
func (c *commonType) Delete(tableName, id string) {
	newId, _ := strconv.ParseInt(id, 10, 64)
	db := service.Sql.DBConnect()
	defer db.Close()
	_, err := db.Exec(`DELETE FROM `+tableName+` WHERE id = ?;`, newId)
	if err != nil {
		panic(err.Error())
	}
}

// 修改 menu_id 指向为空
func (c *commonType) DeleteMenuId(menuId string) {
	db := service.Sql.DBConnect()
	defer db.Close()
	updateTime := time.Now().Unix()
	_, err1 := db.Exec(`UPDATE interface SET update_time = ?, menu_id = NULL WHERE menu_id = ?;`, updateTime, menuId)
	_, err2 := db.Exec(`UPDATE element SET update_time = ?, menu_id = NULL WHERE menu_id = ?;`, updateTime, menuId)
	if err1 != nil {
		panic(err1.Error())
	}
	if err2 != nil {
		panic(err2.Error())
	}
}
