package main

import (
	"traffic-news/common"
	"traffic-news/api"
)

func init() {
	common.ParseCfg()
	common.InitDB()
}

func main() {
	api.RunSrv()
}