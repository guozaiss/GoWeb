package main

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
)

func main() {

	str, err := httplib.Get("http://op.juhe.cn/robot/index?info=你好&key=c9121bfc13b955f01183f082583d1fd6").String()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
	//resp, err := http.Get("http://op.juhe.cn/robot/index?info=你好&key=c9121bfc13b955f01183f082583d1fd6")
	//if err != nil {
	//	fmt.Println("error")
	//}
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//
	//fmt.Println(string(body))
}