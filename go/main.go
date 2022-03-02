package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const VersionSplit = "=android_apk_version="

type VersionDes struct {
	Ver int
	Url string
	Des []string
}

var curVer = &VersionDes{
	Ver: 30201,
	Url: "https://dl.testfairy.com/download/74W3CDHG60RJTC1D82NKRPC60FPAFHT8N9AYAE2D87ZMWQ40/ninja_3.2.00-testfairy.apk",
	Des: []string{
		"1.增加转账功能",
		"2.增加名片功能",
	},
}

func main() {
	result, _ := json.Marshal(curVer)
	fmt.Println(string(result))

	str := fmt.Sprintf("%s%s%s", VersionSplit, result, VersionSplit)
	_ = ioutil.WriteFile("version.md", []byte(str), 0644)
}
