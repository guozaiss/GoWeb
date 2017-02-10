package controllers

import (
	"net/http"
	"fmt"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/httplib"
	"time"
)

func TuLing(w http.ResponseWriter, r *http.Request) {
	DealBegin(w, r)
	bm, _ := cache.NewCache("memory", `{"interval":60}`)
	cache.GetString(bm.Get(r.URL.Path))
	if bm.Get(r.URL.Path) != nil {
		fmt.Fprintf(w, "cache", cache.GetString(bm.Get(r.URL.Path)))
	} else {
		str, _ := httplib.Get("http://op.juhe.cn/robot/index?info=你好&key=c9121bfc13b955f01183f082583d1fd6").String()
		bm.Put(r.URL.Path, str, 10 * time.Second)
		fmt.Fprintf(w, cache.GetString(bm.Get(r.URL.Path)))
	}
	DealEnd(w, r)
}