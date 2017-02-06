package main

import (
	"net/http"
	"fmt"
	"log"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/cache"
	"time"
)

func main() {
	http.HandleFunc("/index", index)

	// 启动web服务，监听9090端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println(r.TLS, r.Host, r.URL.Path, r.Method, "--->", r.Form)
	bm, _ := cache.NewCache("memory", `{"interval":60}`)
	cache.GetString(bm.Get(r.URL.Path))
	if bm.Get(r.URL.Path) != nil {
		fmt.Fprintf(w, "cache", cache.GetString(bm.Get(r.URL.Path)))
	} else {
		str, _ := httplib.Get("http://op.juhe.cn/robot/index?info=你好&key=c9121bfc13b955f01183f082583d1fd6").String()
		bm.Put(r.URL.Path, str, 10 * time.Second)
		fmt.Fprintf(w, cache.GetString(bm.Get(r.URL.Path)),"中国人")
	}
}
