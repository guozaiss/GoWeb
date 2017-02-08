package dbUtils

import (
	"database/sql"
	"fmt"
)

/*
 查询数据库
 */
func Query(sql string) (rows *sql.Rows) {
	mysql := ConnectMysql()
	rows, err := mysql.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	return rows
}

