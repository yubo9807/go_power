package spider

import (
	"server/src/service"
	"time"
)

// 删除某张表中的某条数据
func CommonDelete(tableName, id string) {
	db := service.DBConnect()
	defer db.Close()
	_, err := db.Exec("DELETE FROM ? WHERE id = ?;", tableName, id)
	if err != nil {
		panic(err.Error())
	}
}

// 修改某张表的某一个字段
// 该表中必须包含 update_time 字段（表设计规范约束）
func CommonUpdate(tableName, id, key, value string) {
	db := service.DBConnect()
	defer db.Close()
	updateTime := time.Now().Unix()
	_, err := db.Exec(`UPDATE ? SET update_time = ?, ? = ? WHERE id = ?;`, tableName, updateTime, key, value, id)
	if err != nil {
		panic(err.Error())
	}
}
