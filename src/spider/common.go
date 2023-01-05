package spider

import (
	"server/src/service"
	"strconv"
	"time"
)

// 删除某张表中的某条数据
func CommonDelete(tableName, id string) {
	newId, _ := strconv.ParseInt(id, 10, 64)
	db := service.DBConnect()
	defer db.Close()
	_, err := db.Exec(`DELETE FROM `+tableName+` WHERE id = ?;`, newId)
	if err != nil {
		panic(err.Error())
	}
}

// 修改 menu_id 指向为空
func CommonDeleteMenuId(menuId string) {
	db := service.DBConnect()
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
