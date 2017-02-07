package controllers

import (
	"net/http"
	"fmt"
)

type baseController struct {

}

func DealBegin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("begin")
	fmt.Println(r.TLS, r.Host, r.URL.Path, r.Method, "--->", r.Form, "\n")
}

func DealEnd(w http.ResponseWriter, r *http.Request) {
	fmt.Println("end")
	fmt.Println( "\n")
}