package sys

import (
	"github.com/gin-gonic/gin"
	"sscmgroup.com/app/server/handler/sys"
)

func RegisterLoginController(r *gin.RouterGroup) {
	login := r.Group("/login")
	{
		login.POST("", sys.Login)
		login.GET("/getCaptcha", sys.GetCaptcha)
		login.POST("/loginOut", sys.LoginOut)
	}

}
