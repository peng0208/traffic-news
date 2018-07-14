package collectd

import (
	"time"
	"github.com/PuerkitoBio/goquery"
	"fmt"
)

const (
	// 代理IP站点
	proxySiteUrl = "http://www.xicidaili.com/nn/"
	// 代理IP池更新间隔
	proxyPoolSetInterval = 5 * time.Hour
	// 代理IP更换间隔
	proxyAddrSetInterval = 3 * time.Minute
)

var (
	proxyPool []string
	proxyNum  int
)

func searchProxy() []string {
	hosts := make([]string, 0)

	doc := QueryDoc(proxySiteUrl,"GET")

	doc.Find("#ip_list tbody").Each(func(i int, s *goquery.Selection) {

		s.Find("tr").First().NextAll().Each(func(i int, s *goquery.Selection) {
			ip := s.Find("td").Eq(1).Text()
			port := s.Find("td").Eq(2).Text()
			scheme := s.Find("td").Eq(5).Text()
			host := fmt.Sprintf("%s://%s:%s", scheme, ip, port)
			hosts = append(hosts, host)
		})

	})
	return hosts
}

// 设置代理IP池
func setProxyPool() {
	proxyPool = searchProxy()
	proxyNum = 0
	fmt.Println("更新代理IP池")
}

// 设置代理IP
func setProxyAddr() {
	if len(proxyPool) > 0 {
		proxyNum ++
		fmt.Println("替换代理IP")
	} else {
		fmt.Println("无可用代理IP")
	}
}

// 获取当前使用的代理IP
func GetProxyAddr() string {
	if proxyPool != nil && len(proxyPool) > 0 {
		return proxyPool[proxyNum]
	}
	return ""
}

// 后台任务定时更新代理IP
func proxyTaskSchedule() {
	setProxyPool()
	pt := time.NewTicker(proxyPoolSetInterval)
	at := time.NewTicker(proxyAddrSetInterval)

	for {
		select {
		case <-pt.C:
			setProxyPool()
		case <-at.C:
			setProxyAddr()
		}
	}
}

//func init() {
//	go proxyTaskSchedule()
//}
