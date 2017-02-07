package dbUtils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectMysql() (db *sql.DB) {
	db, err := sql.Open("mysql", "guo:123456@/meizhi")
	if err != nil {
		fmt.Println(err)
	}
	return db
}
