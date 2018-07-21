package api

import (
	"github.com/gin-gonic/gin"
	"os"
	"io"
	"traffic-news/common"
)

func RunSrv() {

	file := common.Cfg.Section("log").Key("log").String()
	host := common.Cfg.Section("server").Key("host").String()
	port := common.Cfg.Section("server").Key("port").String()


	logf, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(logf)
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	addRoutes(r)

	if host == "" {
		host = "127.0.0.1"
	}
	if port == "" {
		port = "8080"
	}
	addr := host + ":" + port

	r.Run(addr)
}

func addRoutes(r *gin.Engine) {
	r.POST("/api/news", newsHandler)
	r.POST("/api/province", provinceHandler)
	r.POST("/api/province/:id", provinceHandler)
	r.POST("/api/code", codeHandler)
	r.POST("/api/code/:id", codeHandler)
}
