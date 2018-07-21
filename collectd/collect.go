package collectd

import (
	"time"
	"regexp"
	"github.com/PuerkitoBio/goquery"
	"traffic-news/common"
	"strings"
)

const baseUrl = "https://www.icauto.com.cn/gonglu/"

type (
	news struct {
		time    string
		content string
	}

	provinceNews struct {
		time     string
		content  string
		province string
	}

	codeNews struct {
		time    string
		content string
		code    string
		name    string
	}

	province struct {
		name string
		url  string
	}

	code struct {
		code string
		name string
		url  string
	}

	newsList []*news
	provinceNewsList []*provinceNews
	codeNewsList []*codeNews
	provinceList []*province
	codeList []*code
)

func searchIndex() *goquery.Document {
	doc := QueryDoc(baseUrl, "GET")
	return doc
}

func searchNewsList(doc *goquery.Document) newsList {
	data := make(newsList, 0)
	re := regexp.MustCompile(`^[\s]*([\d]{4}-[\d]{2}-[\d]{2}[\s]*[\d]{2}:[\d]{2}).{2}(.+)$`)

	listAll := doc.Find("div #ScrollMe").Children()
	for i := 0; i < 10; i++ {
		li := listAll.Eq(i).Text()
		r := re.FindStringSubmatch(li)
		if len(r) > 0 && len(r[2]) > 0 {
			data = append(data, &news{r[1], replace(r[2])})
		}
	}
	return data
}

func searchProvList(doc *goquery.Document) provinceList {
	data := make(provinceList, 0)
	re := regexp.MustCompile(`^[\s]*(.+)高速公路$`)

	doc.Find("div .articlemain .thethirdB ul li").Each(func(i int, s *goquery.Selection) {
		url, _ := s.Find("a").Attr("href")
		title := s.Text()
		r := re.FindStringSubmatch(title)
		if len(r) > 0 && len(r[1]) > 0 {
			data = append(data, &province{r[1] + "省", url})
		}
	})
	return data
}

func searchCodeList(doc *goquery.Document) codeList {
	data := make(codeList, 0)
	re := regexp.MustCompile(`^[\s]*(.+)\((.+)\)$`)

	doc.Find("div .articlemain .thethirdB ul li").Each(func(i int, s *goquery.Selection) {
		url, _ := s.Find("a").Attr("href")
		title := s.Text()
		r := re.FindStringSubmatch(title)
		if len(r) > 0 {
			data = append(data, &code{r[2], r[1], url})
		}
	})

	return data
}

func (pl provinceList) search() {
	for _, p := range pl {
		provNewsData := make(provinceNewsList, 0)
		time.Sleep(requestInterval)

		doc := QueryDoc(p.url, "GET")
		doc.Find("div .lk-body-main .lknew p").Each(func(i int, s *goquery.Selection) {
			span := s.Find("span")
			span.Eq(0).Children().Remove()
			tm := span.Eq(0).Text()
			content := span.Eq(1).Text()
			provNewsData = append(provNewsData, &provinceNews{tm, replace(content), p.name})
		})
		provNewsData.save()
	}
}

func (cl codeList) search() {
	for _, c := range cl {
		codeNewsData := make(codeNewsList, 0)
		time.Sleep(requestInterval)

		doc := QueryDoc(c.url, "GET")
		doc.Find("div .lk-body-main .lknew p").Each(func(i int, s *goquery.Selection) {
			span := s.Find("span")
			span.Eq(0).Children().Remove()
			tm := span.Eq(0).Text()
			content := span.Eq(1).Text()
			codeNewsData = append(codeNewsData, &codeNews{tm, replace(content), c.code, c.name})
		})
		codeNewsData.save()
	}

}

func (il newsList) save() {
	for _, i := range il {
		SaveNews(i.time, i.content)
	}
}

func (pl provinceList) save() {
	for _, i := range pl {
		SaveProvince(i.name)
	}
}

func (cl codeList) save() {
	for _, i := range cl {
		SaveCode(i.code, i.name)
	}
}

func (pnl provinceNewsList) save() {
	for _, pn := range pnl {
		SaveProvinceNews(pn.time, pn.province, pn.content)
	}
}

func (cnl codeNewsList) save() {
	for _, cn := range cnl {
		SaveCodeNews(cn.time, cn.code, cn.content)
	}
}

func checkTime() {
	hour := time.Now().Hour()
	switch {
	case hour > 0 && hour < 6:
		common.Logger().Info("当前时间段不允许采集,暂停线程")
		zZzZ := time.Duration(6 - hour)
		time.Sleep(zZzZ * time.Hour)
	}
}

// 去除多余空白，多余字符
func replace(s string) string {
	str := strings.Replace(s, " ","", -1)
	str = strings.Replace(s, "r","", -1)
	return str
}

func collectTask() {
	doc := searchIndex()

	news := searchNewsList(doc)
	provList := searchProvList(doc)
	codeList := searchCodeList(doc)

	news.save()
	provList.save()
	codeList.save()

	provList.search()
	codeList.search()
}

var (
	requestInterval  time.Duration
	scheduleInterval time.Duration
	taskCount            int64
)

func CollectTaskSchedule() {
	ri, _ := common.Cfg.Section("server").Key("request_interval").Int64()
	ti, _ := common.Cfg.Section("server").Key("task_interval").Int64()

	requestInterval = time.Duration(ri) * time.Second
	scheduleInterval = time.Duration(ti) * time.Second

	common.Logger().Info("开启数据采集线程")

	c := make(chan struct{}, 1)
	t := time.NewTicker(scheduleInterval)

	go func() {
		for {
			select {
			case <-t.C:
				<-c
			}
		}
	}()

	for {
		c <- struct{}{}

		checkTime()
		counter()
		collectTask()
	}
}

func counter() {
	taskCount ++
	common.Logger().Infof("开始第 %d 次采集", taskCount)
}