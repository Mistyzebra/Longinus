package web

import (
	"github.com/gin-gonic/gin"
)

func RunServer() {
	router := gin.Default()
	router.POST("/longinuscli/cwe22", auditCWE22)
	router.Run(":8080") // 在8080端口上运行
}

func auditCWE22(c *gin.Context) {
	// 实现审计CWE-22逻辑
}
