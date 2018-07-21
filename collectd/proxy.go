package collectd

import (
	"traffic-news/common"
	"fmt"
	"net/http"
	"io/ioutil"
)

func GetProxy() string {
	return GetProxyAddr(GetMoguApi())
}

func GetProxyAddr(url string) string {
	if url == "" {
		return ""
	}

	resp, err := http.Get(url)
	if err != nil {
		common.Logger().Error(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		common.Logger().Error(err)
	}

	addr := "http://%s" + string(body)
	return addr
}

func GetMoguApi() string {
	var api string
	sql := "select key, count from t_key limit 1;"
	result, _ := common.Query(sql)
	if len(result) > 0 {
		api = fmt.Sprintf(
			"http://piping.mogumiao.com/proxy/api/get_ip_bs?appKey=%s&count=%s&expiryDate=0&format=2&newLine=2",
			result[0]["key"],
			result[0]["count"],
		)
	}
	return api
}
