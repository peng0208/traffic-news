package main

import (
	"traffic-news/common"
	"traffic-news/collectd"
)

func init() {
	common.ParseCfg()
	common.InitDB()
}

func main() {
	go collectd.CollectTaskSchedule()
	select {}
}
