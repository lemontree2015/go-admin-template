package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func NoCache(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
	c.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
	c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	c.Next()
}

func Options(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
	//允许类型校验
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, "ok!")
	}
	c.Next()
}

func BanEdit(c *gin.Context) {
	if c.Request.URL.Path != "/login" && c.Request.Method != "GET" {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusForbidden,
			"msg":  "为了更好体验，禁止修改数据^-^",
		})
	}
	c.Next()
}
