package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"../controllers"
	"fmt"
)

func init() {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.Index)
	router.HandleFunc("/index", controllers.Index)
	router.HandleFunc("/queryPerson", controllers.QueryPerson)
	router.HandleFunc("/tuLing", controllers.TuLing)
	http.Handle("/", router)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("err==>", err)
	}
}
