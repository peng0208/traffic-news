package main

import (
	"traffic-news/common"
	"traffic-news/api"
)

func init() {
	common.ParseCfg()
	common.InitDB()
	common.InitRedis()
}

func main() {
	api.RunSrv()
}