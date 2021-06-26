package middleware

import "github.com/gin-gonic/gin"

func CommonMiddleware(r *gin.Engine) {
	// NoCache is a middleware function that appends headers
	r.Use(NoCache)
	// 跨域处理
	r.Use(Options)
	// 禁止修改数据
	r.Use(BanEdit)
}
