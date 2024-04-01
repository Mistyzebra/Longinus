package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化 Gin 框架
	r := gin.Default()

	// 设置路由
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	// 运行 Gin 框架
	r.Run(":8080")
}
