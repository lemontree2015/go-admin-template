package sys

import (
	"github.com/gin-gonic/gin"
	"net/http"
	dto "sscmgroup.com/app/dto/sys"
	"sscmgroup.com/app/logger"
	"strconv"
)

func AddMenu(c *gin.Context) {
	var menuForm dto.MenuForm
	err := c.Bind(&menuForm)
	if err != nil {
		logger.Logger.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}

	menu, err := srv.AddMenu(menuForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": menu,
		"msg":  "ok",
	})
}

func GetMenu(c *gin.Context) {
	Id := c.Param("id")
	menuId, _ := strconv.Atoi(Id)
	menu, err := srv.GetMenuById(menuId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": menu,
		"msg":  "ok",
	})
}

func MenusList(c *gin.Context) {
	list, err := srv.GetAllSysMenus()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": list,
		"msg":  "ok",
	})
}

func EditMenu(c *gin.Context) {
	var uf dto.MenuForm
	err := c.Bind(&uf)
	if err != nil {
		logger.Logger.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}

	menu, err := srv.GetMenuById(uf.MenuId)
	if menu == nil || err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	menu, err = srv.UpdateMenu(menu, uf)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": menu,
		"msg":  "ok",
	})
}
