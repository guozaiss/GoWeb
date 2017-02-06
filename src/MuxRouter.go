package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/products", ProductsHandler)
	router.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", router)
	_ = http.ListenAndServe(":9090", nil)
}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HomeHandler")
}
func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ProductsHandler")
}
func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ArticlesHandler")
}