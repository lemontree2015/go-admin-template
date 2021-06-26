package sys

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"sscmgroup.com/app/config"
	dto "sscmgroup.com/app/dto/sys"
	"sscmgroup.com/app/logger"
	"sscmgroup.com/app/model/sys"
	"sscmgroup.com/app/module/captcha"
	"time"
)

func Login(c *gin.Context) {
	var loginForm dto.LoginForm
	err := c.Bind(&loginForm)
	if err != nil {
		logger.Logger.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}

	if !captcha.Verify(loginForm.UUID, loginForm.Code, true) {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "验证码错误",
		})
		return
	}

	logger.Logger.Error(c.Request.Header.Get("X-Real-IP"), ":", c.ClientIP())
	u, err := srv.Login(loginForm, c.Request.Header.Get("X-Real-IP"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}

	token, expires, err := buildToken(u)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code":   200,
		"token":  token,
		"expire": expires,
		"msg":    "ok",
	})
}

func LoginOut(c *gin.Context) {
	c.SetCookie("auth", "", -1, "", "", false, true)
	c.SetCookie("userId", "", -1, "", "", false, false)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
	})
}

func GetCaptcha(c *gin.Context) {
	id, b64s, err := captcha.DriverDigitFunc()
	if err != nil {
		logger.Logger.Error("DriverDigitFunc error", err.Error())
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "验证码获取失败",
		})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": 200,
		"data": b64s,
		"id":   id,
		"msg":  "ok",
	})
}

func buildToken(user *sys.User) (token string, expires time.Time, err error) {
	if user == nil {
		err = errors.New("用户信息不能为空")
		return
	}
	jwtConf := config.Conf.Jwt["mgr"]
	fmt.Println("jwtConf...",jwtConf)
	token, expires, err = buildJwt(user, jwtConf)
	if err != nil {
		logger.Logger.Error("签名生成失败: %s", err.Error())
		return
	}
	return
}
func buildJwt(user *sys.User, jwtConf config.JwtConfig) (jwtStr string, expires time.Time, err error) {
	now := time.Now()
	var roles []string
	if len(user.Roles) > 0 {
		for _, r := range user.Roles {
			roles = append(roles, r.Name)
		}
	}
	userClaim := dto.UserJwt{Id: user.Id, UserName: user.UserName, Roles: roles}
	userClaim.Issuer = "mgr.sscmgroup.com"
	userClaim.ExpiresAt = now.Unix() + int64(jwtConf.TTL)
	userClaim.IssuedAt = now.Unix()
	expires = now.Add(time.Duration(jwtConf.TTL))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim)
	jwtStr, err = token.SignedString([]byte(jwtConf.Secret))
	return
}
