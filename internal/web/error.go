package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Error404 错误页面
func Error404(c *gin.Context) {
	c.HTML(http.StatusOK, "error/404", gin.H{
		"title": "404未找到",
		"ctx":   c,
	})
}
