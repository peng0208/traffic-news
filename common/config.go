package common

import (
	"github.com/go-ini/ini"
	"flag"
)

var file string
var Cfg *ini.File

func parseArg() {
	flag.StringVar(&file, "c", "config.ini", "configfile 配置文件")
	flag.Parse()
}

func ParseCfg() {
	parseArg()
	cfg, er := ini.Load(file)
	CheckError(er)
	Cfg = cfg
}

func GetCfg() *ini.File {
	cfg, er := ini.Load(file)
	CheckError(er)
	return cfg
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
