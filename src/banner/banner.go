package banner

import (
	"fmt"
	. "configuration"
	. "github.com/dimiro1/banner"
	"github.com/mattn/go-colorable"
	"bytes"
	"io/ioutil"
)

func LoadBanner() {
	var banner = Config.AppDesc
	file, e := ioutil.ReadFile("src/banner/banner.txt")
	if e == nil { banner = string(file) }

	Init(colorable.NewColorableStdout(), true, true, bytes.NewBufferString(banner))

	fmt.Println("[33m[REDIS]:")
	fmt.Println("\t[33m[HOST]:", Config.Redis.Host + ":" + Config.Redis.Port)
	fmt.Println("\t[33m[DB]:", Config.Redis.DB)
	if Config.Redis.Password != "" {
		fmt.Println("\t[33m[PASS]:", Config.Redis.Password)
	}
}
