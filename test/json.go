package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonBlob = []byte(`{"sounds/ambient/cave/cave1.ogg": {"hash": "fd6f6042f2a4ca8011e16c5199221c16db897f26","size": 1876295}}`)
	type FileInfo struct {
		Hash string
		Size int64
	}

	var file map[string]FileInfo

	err := json.Unmarshal(jsonBlob, &file)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", file)
}
