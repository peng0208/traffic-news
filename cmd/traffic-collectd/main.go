package main

import (
	"traffic-news/common"
	"traffic-news/collectd"
	"fmt"
)

func init() {
	common.ParseCfg()
	common.InitDB()
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	go collectd.CollectTaskSchedule()

	select {}
}
