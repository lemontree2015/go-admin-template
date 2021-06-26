package sys

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sscmgroup.com/app/casbin"
	"sscmgroup.com/app/common"
	"sscmgroup.com/app/logger"
)

func CasbinMiddleware(c *gin.Context) {
	p := c.Request.URL.Path
	m := c.Request.Method
	user := AuthInfo(c)
	logger.Logger.Debug(user.User.UserName, p, m)
	b, err := casbin.Enforcer.Enforce(user.User.UserName, p, m)
	if common.InArray(common.AdminTag, user.User.Roles) {
		b = true
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusForbidden,
			"msg":  err.Error(),
		})
	} else if !b {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusForbidden,
			"msg":  "你没有" + p + "接口权限，请联系管理员",
		})
	}
	c.Next()
}
