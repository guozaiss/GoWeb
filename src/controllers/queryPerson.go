package controllers

import (
	"net/http"
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"../dbUtils"
	"../models"
)

type queryPersonController struct {

}

func QueryPerson(w http.ResponseWriter, r *http.Request) {
	DealBegin(w, r)
	//rows := dbUtils.Query("select * from showalltb_meizhi")
	engine, err := xorm.NewEngine("mysql", "guo:123456@/meizhi")
	if err!=nil {
		fmt.Println(err)
	}
	//engine.Sync2(new(models.Meizhi))

	ss,err :=engine.Query("select * from tb_meizhi where id ='id_01'")
	fmt.Println(ss)
	status,err :=engine.Insert(models.Meizhi{Id:"ididid",FileName:"FilefileName"})
	fmt.Println(status)

	rows := dbUtils.Query("select * from tb_meizhi where id ='id_01'")
	for rows.Next() {
		var id, fileName string
		rows.Columns()
		err := rows.Scan(&id, &fileName)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintf(w,"id", id, "fileName", fileName)
	}
	DealEnd(w, r)
}
