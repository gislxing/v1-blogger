package main

import (
	"github.com/gin-gonic/gin"
	_ "zbs/v1-blogger/dao"
)

func main() {
	r := gin.Default()

	// 设置以 debug 模式运行gin
	// 生产中使用 gin.ReleaseMode 模式
	gin.SetMode(gin.DebugMode)

	r.Run(":9000")
}
