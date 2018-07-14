package collectd

import (
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"time"
	"fmt"
)

const (
	requestTimeout = 30 * time.Second
	userAgent      = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36"
	referer        = "https://www.icauto.com.cn/gonglu/"
	acceptLanguage = "zh-CN,zh;q=0.9,en;q=0.8"
	accept         = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8"
	connection     = "keep-alive"
	cacheControl   = "max-age=0"
)

var CookiesCache []*http.Cookie

func QueryDoc(urlstr string, method string) *goquery.Document {
	/*
	var clt *http.Client
	var trans *http.Transport
	proxyAddr := GetProxyAddr()
	if proxyAddr == "" {
		clt = &http.Client{Timeout: requestTimeout}
	} else {
		proxy := func(_ *http.Request) (*url.URL, error) {
			return url.Parse(proxyAddr)
		}
		trans = &http.Transport{Proxy: proxy}
		clt = &http.Client{Timeout: requestTimeout, Transport: trans}
	}
	*/

	clt := &http.Client{Timeout: requestTimeout}
	req, _ := http.NewRequest(method, urlstr, nil)
	if CookiesCache != nil {
		for _, c := range CookiesCache {
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
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()
	ret := res.StatusCode
	if ret != 200 {
		return nil
	}
	CookiesCache = res.Cookies()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	return doc
}
