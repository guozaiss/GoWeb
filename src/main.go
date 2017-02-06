package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	resp, err := http.Get("http://op.juhe.cn/robot/index?info=你好&key=c9121bfc13b955f01183f082583d1fd6")
	if err != nil {
		fmt.Println("error")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}

