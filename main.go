package ninjahome_github_io

import (
	"encoding/json"
	"fmt"
	"github.com/ninjahome/ninjahome.github.io/golib"
	"io/ioutil"
)

func main() {
	result, _ := json.Marshal(golib.CurVer)
	fmt.Println(string(result))

	str := fmt.Sprintf("%s%s%s", golib.VersionSplit, result, golib.VersionSplit)
	_ = ioutil.WriteFile("version.md", []byte(str), 0644)
}
