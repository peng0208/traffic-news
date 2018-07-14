package common

import (
	"github.com/go-ini/ini"
	"flag"
	"fmt"
)

var file string
var Cfg *ini.File

func parseArg() {
	flag.StringVar(&file, "c", `config.ini`, "configfile")
	flag.Parse()
}

func ParseCfg() {
	fmt.Println("加载配置文件")
	parseArg()
	cfg, er := ini.Load(file)
	CheckError(er)
	Cfg = cfg
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
