package controllers

import (
	"net/http"
	"fmt"
)
type baseController struct {

}
func Deal(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println(r.TLS, r.Host, r.URL.Path, r.Method, "--->", r.Form)
}