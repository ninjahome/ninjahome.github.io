package main

import (
	"encoding/json"
	"fmt"
)

type VersionDes struct {
	Ver int
	Des []string
}

var curVer = &VersionDes{
	Ver: 30201,
	Des: []string{
		"1.增加转账功能",
		"2.增加名片功能",
	},
}

func main() {
	result, _ := json.Marshal(curVer)
	fmt.Println(string(result))
}
