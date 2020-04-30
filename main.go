package main

import (
	"ParsesHtml/src"
	"github.com/gin-gonic/gin"
	"log"
)

func main(){
	//create route
	router := gin.Default()
	gin.SetMode(gin.DebugMode)
	//设置html目录 和img目录
	router.LoadHTMLGlob("html/*")
	router.Static("/img", "img") // 第一个参数表示http请求的路径
	router.Static("/js","js")
	//使用跨域中间件
	router.Use(src.Cors())

	//为路由添加访问路径
	router.GET("/", src.Index)
	router.POST("/parse", src.ParseHtml)
	if err := router.Run(":81");err != nil{
		log.Fatal(err.Error())
	}
}

