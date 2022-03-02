package golib

const VersionSplit = "=android_apk_version="

type VersionDes struct {
	Ver int
	Url string
	Des []string
}

var CurVer = &VersionDes{
	Ver: 30201,
	Url: "https://dl.testfairy.com/download/74W3CDHG60RJTC1D82NKRPC60FPAFHT8N9AYAE2D87ZMWQ40/ninja_3.2.00-testfairy.apk",
	Des: []string{
		"1.增加转账功能",
		"2.增加名片功能",
	},
}
