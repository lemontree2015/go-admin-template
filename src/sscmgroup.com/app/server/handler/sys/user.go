package sys

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"sscmgroup.com/app/config"
	dto "sscmgroup.com/app/dto/sys"
	"sscmgroup.com/app/logger"
	"strconv"
)

func GetProfile(c *gin.Context) {
	AuthId, ok := c.Get("_authId")
	if !ok {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "当前未登陆",
		})
		return
	}
	u, err := srv.GetUserById(AuthId.(int))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": u,
		"msg":  "success",
	})
}

func AccessSideBar(c *gin.Context) {
	AuthId, ok := c.Get("_authId")
	if !ok {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "当前未登陆",
		})
		return
	}
	menus, err := srv.GetUserAccessSideBar(AuthId.(int))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": 200,
		"data": menus,
		"msg":  "success",
	})
}

func AddUser(c *gin.Context) {
	var userForm dto.UserForm
	err := c.Bind(&userForm)
	if err != nil {
		logger.Logger.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}

	u, err := srv.AddUser(userForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": u,
		"msg":  "ok",
	})
}

func GetUser(c *gin.Context) {
	Id := c.Param("id")
	userId, _ := strconv.Atoi(Id)
	u, err := srv.GetUserById(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": u,
		"msg":  "ok",
	})
}

func EditUser(c *gin.Context) {
	var uf dto.UserForm
	err := c.Bind(&uf)
	if err != nil {
		logger.Logger.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}

	user, err := srv.GetUserById(uf.UserId)
	if user == nil || err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	user, err = srv.UpdateUser(user, uf)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": user,
		"msg":  "ok",
	})
}

func DelUser(c *gin.Context) {
	Id := c.Param("id")
	userId, _ := strconv.Atoi(Id)
	u, err := srv.GetUserById(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	u.Status = 0
	err = srv.UpdateUserInfo(u)
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": u,
		"msg":  "ok",
	})
}

func UserList(c *gin.Context) {
	var us dto.SysUserSearch
	err := c.Bind(&us)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}

	count, list, err := srv.GetSysUserPage(&us)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": map[string]interface{}{
			"list":      list,
			"count":     count,
			"pageIndex": us.PageIndex,
			"PageSize":  us.PageSize,
		},
		"msg": "ok",
	})
	return
}

func ChangePwd(c *gin.Context) {
	var uf dto.ChangePwdForm
	err := c.Bind(&uf)
	if err != nil {
		logger.Logger.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}

	AuthId, ok := c.Get("_authId")
	if !ok {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "当前未登陆",
		})
		return
	}
	u, err := srv.GetUserById(AuthId.(int))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  err.Error(),
		})
		return
	}

	err = srv.ChangeUserPwd(u, uf.OldPassword, uf.NewPassword)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
	})
}

func UploadAvatar(c *gin.Context) {
	AuthId, ok := c.Get("_authId")
	if !ok {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "当前未登陆",
		})
		return
	}
	u, err := srv.GetUserById(AuthId.(int))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  err.Error(),
		})
		return
	}
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	uploadCfg := config.Conf.Uploads["avatar"]
	fileName := strconv.Itoa(u.Id) + ".jpg"
	filPath := uploadCfg.Dir + fileName
	for _, file := range files {
		log.Debugf("upload avatar file: %s", file.Filename)
		// 上传文件至指定目录
		err := c.SaveUploadedFile(file, filPath)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
			})
			return
		}
	}

	//u.Avatar = "/" + filPath
	u.Avatar = uploadCfg.CdnHost + fileName
	err = srv.UpdateUserInfo(u)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"avatar": uploadCfg.CdnHost + fileName,
		"msg":    "ok",
	})
}
