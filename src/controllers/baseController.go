package controllers

import (
	"net/http"
	"fmt"
	"time"
)

type baseController struct {

}

func DealBegin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("begin",time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(r.TLS, r.Host, r.URL.Path, r.Method, "--->", r.Form)
}

func DealEnd(w http.ResponseWriter, r *http.Request) {
	fmt.Println("end","\n")
}