package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Index 首页
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index", gin.H{
		"title": "首页",
		"ctx":   c,
	})
}
