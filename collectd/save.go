package collectd

import (
	"traffic-news/common"
	"fmt"
)

func SaveNews(params ...interface{}) {
	sql := "INSERT IGNORE INTO `t_news` (`time`, `content`) VALUES (?,?);"
	if err := common.Execute(sql, params...); err != nil {
		fmt.Println(err)
	}
}
func SaveProvince(params interface{}) {
	sql := "INSERT IGNORE INTO `t_province` (`name`) VALUES (?);"
	if err := common.Execute(sql, params); err != nil {
		fmt.Println(err)
	}
}
func SaveCode(params ...interface{}) {
	sql := "INSERT IGNORE INTO `t_code` (`code`, `name`) VALUES (?,?);"
	if err := common.Execute(sql, params...); err != nil {
		fmt.Println(err)
	}
}
func SaveProvinceNews(params ...interface{}) {
	sql := "INSERT IGNORE INTO `t_province_news` (`time`, `province`, `content`) VALUES (?,?,?);"
	if err := common.Execute(sql, params...); err != nil {
		fmt.Println(err)
	}
}
func SaveCodeNews(params ...interface{}) {
	sql := "INSERT IGNORE INTO `t_code_news` (`time`, `code`, `content`) VALUES (?,?,?);"
	if err := common.Execute(sql, params...); err != nil {
		fmt.Println(err)
	}
}
