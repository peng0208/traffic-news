package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func newsHandler(c *gin.Context) {
	data := getNews()
	c.JSON(http.StatusOK, data)
}

func provinceHandler(c *gin.Context) {
	provId := c.Param("id")
	if provId == "" {
		data := getProvince()
		c.JSON(http.StatusOK, data)
	} else {
		data := getProvinceNews(provId)
		c.JSON(http.StatusOK, data)
	}
}

func codeHandler(c *gin.Context) {
	codeId := c.Param("id")
	if codeId == "" {
		data := getCode()
		c.JSON(http.StatusOK, data)
	} else {
		data := getCodeNews(codeId)
		c.JSON(http.StatusOK, data)
	}
}

// 登录, wxCode, openid, appSecret, 获取session
func loginHandler(c *gin.Context) {
	wxCode := c.Query("code")
	session := getSessionId(c, wxCode)
	if session != nil && setSession(session.sessionId, session.sessionValue) {
		c.String(http.StatusOK, session.sessionId)
		return
	}
	c.String(http.StatusUnauthorized,"401")
}
