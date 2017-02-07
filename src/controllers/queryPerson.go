package controllers

import (
	"net/http"
	"fmt"
	"../dbUtils"
)

type queryPersonController struct {

}

func QueryPerson(w http.ResponseWriter, r *http.Request) {
	DealBegin(w, r)
	rows := dbUtils.Query("select * from showalltb_meizhi")
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
