package sys

import (
	"github.com/gin-gonic/gin"
	"net/http"
	dto "sscmgroup.com/app/dto/sys"
	"sscmgroup.com/app/logger"
	"strconv"
)

func AddSetting(c *gin.Context) {
	var form dto.SettingForm
	err := c.Bind(&form)
	if err != nil {
		logger.Logger.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}

	row, err := srv.AddSetting(form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": row,
		"msg":  "ok",
	})
}

func SettingList(c *gin.Context) {
	var search dto.SettingSearch
	err := c.Bind(&search)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}

	count, list, err := srv.GetSettingByPage(&search)
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
			"pageIndex": search.PageIndex,
			"PageSize":  search.PageSize,
		},
		"msg": "ok",
	})
	return
}

func GetSetting(c *gin.Context) {
	Id := c.Param("id")
	settingId, _ := strconv.Atoi(Id)
	row, err := srv.GetSettingById(settingId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": row,
		"msg":  "ok",
	})
}

func EditSetting(c *gin.Context) {
	var form dto.SettingForm
	err := c.Bind(&form)
	if err != nil {
		logger.Logger.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}

	row, err := srv.GetSettingById(form.Id)
	if row == nil || err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	row, err = srv.UpdateSetting(row, form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": row,
		"msg":  "ok",
	})
}
