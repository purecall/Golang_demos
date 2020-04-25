package main

import (
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
)

func initTemplate(router *gin.Engine) {
	router.StaticFile("/", "./static/index.html")
	router.StaticFile("/favicon.ico", "./static/favicon.ico")
	router.Static("/css/", "./static/css/")
	router.Static("/fonts/", "./static/fonts/")
	router.Static("/img/", "./static/img/")
	router.Static("/js/", "./static/js/")
}

func main() {
	router := gin.Default()

	ginpprof.Wrapper(router)
	initTemplate(router)

	router.Run(":8080")
}
