package collectd

import (
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"time"
	"net/url"
	"traffic-news/common"
)

const (
	requestTimeout = 15 * time.Second
	userAgent      = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36"
	referer        = "https://www.icauto.com.cn/gonglu/"
	acceptLanguage = "zh-CN,zh;q=0.9,en;q=0.8"
	accept         = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8"
	connection     = "keep-alive"
	cacheControl   = "max-age=0"
)

var cookiesCache []*http.Cookie

func requestDoc(urlstr string, method string, proxy string) *goquery.Document {
	var clt *http.Client
	var trans *http.Transport

	if proxy == "" {
		clt = &http.Client{Timeout: requestTimeout}
	} else {
		prx := func(_ *http.Request) (*url.URL, error) {
			return url.Parse(proxy)
		}
		trans = &http.Transport{Proxy: prx}
		clt = &http.Client{Timeout: requestTimeout, Transport: trans}
	}

	req, _ := http.NewRequest(method, urlstr, nil)
	if cookiesCache != nil {
		for _, c := range cookiesCache {
			req.AddCookie(c)
		}
	}
	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("Referer", referer)
	req.Header.Add("Accept-Language", acceptLanguage)
	req.Header.Add("Accept", accept)
	req.Header.Add("Connection", connection)
	req.Header.Add("Cache-Control", cacheControl)

	res, err := clt.Do(req)
	if err != nil {
		common.Logger().Error(err)
		return nil
	}
	defer res.Body.Close()
	cookiesCache = res.Cookies()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		common.Logger().Error(err)
		return nil
	}
	return doc
}

var currentProxy string

func QueryDoc(urlstr string, method string) *goquery.Document {
	var doc *goquery.Document

	for {
		if currentProxy == "" {
			currentProxy = GetProxy()
			common.Logger().Infof("获取代理IP: [%s]", currentProxy)

		}
		doc = requestDoc(urlstr, method, currentProxy)

		if doc == nil {
			doc = requestDoc(urlstr, method, currentProxy)

			if doc == nil {
				common.Logger().Errorf("请求异常或代理IP失效: [%s], 1秒后重试", currentProxy)
				// 删除无效代理IP
				currentProxy = ""
				time.Sleep(1 * time.Second)
				continue
			}
			return doc
		}
		return doc

	}

}
