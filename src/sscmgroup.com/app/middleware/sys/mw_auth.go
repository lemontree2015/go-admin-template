package sys

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"sscmgroup.com/app/config"
	dto "sscmgroup.com/app/dto/sys"
)

func AuthInfo(c *gin.Context) *dto.Auth {
	return c.MustGet("_auth").(*dto.Auth)
}

// UserAuthMiddleware 用户授权中间件
func AuthorizedMiddle(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	if tokenString == "" {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "当前未登陆",
		})
		return
	}

	jwtConf := config.Conf.Jwt["mgr"]
	token, err := jwt.ParseWithClaims(tokenString, &dto.UserJwt{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(jwtConf.Secret), nil
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  err.Error(),
		})
		return
	}

	auth := &dto.Auth{IsAuth: false}
	if claim, ok := token.Claims.(*dto.UserJwt); ok && token.Valid {
		auth.IsAuth = true
		auth.User = dto.AuthUser{Id: claim.Id, UserName: claim.UserName, Roles: claim.Roles}
		c.Set("_auth", auth)
		c.Set("_authId", auth.User.Id)
	}
	if auth.IsAuth == false {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "当前未登陆",
		})
	}
	c.Next()
}
