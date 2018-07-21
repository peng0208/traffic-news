package api

import (
	"traffic-news/common"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"crypto/md5"
	"encoding/hex"
)

var (
	appId     = "wx61a706ec6e0fb831"
	appSecret = "49d0e39da312ab78b11b7faf21505b43"
	wxAuthUrl = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
)

func getNews() []map[string]string {
	sql := `select DATE_FORMAT(time,'%Y-%m-%d %H:%i') as time , content from t_news order by time desc limit 30;`
	data, _ := common.Query(sql)
	return data
}

func getProvince() []map[string]string {
	sql := `select id, name from t_province;`
	data, _ := common.Query(sql)
	return data
}

func getCode() []map[string]string {
	sql := `select id, name, code from t_code;`
	data, _ := common.Query(sql)
	return data
}

func getProvinceNews(province string) []map[string]string {
	sql := `select DATE_FORMAT(time,'%Y-%m-%d %H:%i') as time, content from t_province_news n join t_province p 
			on n.province = p.name WHERE p.id = ? order by time desc limit 30;`
	data, _ := common.Query(sql, province)
	return data
}

func getCodeNews(code string) []map[string]string {
	sql := `select DATE_FORMAT(time,'%Y-%m-%d %H:%i') as time, content from t_code_news n join t_code c 
			on n.code = c.code WHERE c.id = ? order by time desc limit 30;`
	data, _ := common.Query(sql, code)
	return data
}

type sessionData struct {
	sessionId    string
	sessionValue string
}

func getSessionId(c *gin.Context, wxCode string) *sessionData {
	result := make(map[string]interface{})
	wxUrl := fmt.Sprintf(wxAuthUrl, appId, appSecret, wxCode)
	res, _ := http.Get(wxUrl)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	if err := json.Unmarshal(body, &result); err != nil {
		return nil
	}
	if _, ok := result["session_key"]; ok {
		openid := result["openid"].(string)
		key := result["session_key"].(string)
		sessionId := md5Hash(openid, key)
		return &sessionData{sessionId: sessionId, sessionValue: key}
	} else {
		return nil
	}
}

func md5Hash(s1, s2 string) string {
	h := md5.New()
	h.Write([]byte(s1 + s2))
	return hex.EncodeToString(h.Sum(nil))
}


func setSession(k,v string) bool {
	return common.SetWithTTL(k, v,"1800")
}

func getSession(k string) string {
	return common.Get(k)
}