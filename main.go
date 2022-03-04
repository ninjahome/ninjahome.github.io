package main

import (
	"encoding/json"
	"fmt"
	"github.com/ninjahome/ninjahome.github.io/golib"
	"io/ioutil"
)

var VersionValue = &golib.VersionDes{
	Ver:     30202,
	VerName: "3.2.02",
	Url:     "https://ninjahome.github.io/ninja.apk",
	Des: []string{
		"1.优化消息处理流程",
		"2.增加名片功能",
	},
}

func main() {
	result, _ := json.Marshal(VersionValue)
	fmt.Println(string(result))

	_ = ioutil.WriteFile("version.js", result, 0644)
}
