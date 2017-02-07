package controllers

import (
	"net/http"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type queryPersonController struct {

}

func QueryPerson(w http.ResponseWriter, r *http.Request) {
	Deal(w, r)
	r.ParseForm()
	i := r.Form.Get("userId")
	fmt.Println(i)
	db, err := sql.Open("mysql", "guo:123456@/meizhi")
	if err != nil {
		fmt.Println(err)
	}
	rows, err := db.Query("select * from showalltb_meizhi")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var id, fileName string
		rows.Columns()
		err := rows.Scan(&id, &fileName)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("id", id, "fileName", fileName)
	}

}
