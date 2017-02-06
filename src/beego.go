package main

import (
	"github.com/astaxie/beego"
	"net/http"
)

func main() {
	http.HandleFunc("/index", beego.Controller{})
}
