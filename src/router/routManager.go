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
	//err := http.ListenAndServeTLS(":9090","c:/Users/Admin/IdeaProjects/Web/src/router/server.crt",
	//	"c:/Users/Admin/IdeaProjects/Web/src/router/server.key", nil)
	if err != nil {
		fmt.Println("err==>", err)
	}
}
