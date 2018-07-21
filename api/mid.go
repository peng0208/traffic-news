package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 登录态,验证session
func authRequired() gin.HandlerFunc {

	return func(c *gin.Context) {
		if c.Request.URL.Path != "/api/session" {
			sid := c.Query("sid")
			v := getSession(sid)
			if v == "" {
				c.String(http.StatusUnauthorized, "401")
			}
			setSession(sid, v)
		}
	}
}
