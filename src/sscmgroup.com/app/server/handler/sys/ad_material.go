package sys

import (
	"github.com/gin-gonic/gin"
	"net/http"
	dto "sscmgroup.com/app/dto/sys"
	"sscmgroup.com/app/logger"
	"strconv"
)

func AddMaterial(c *gin.Context) {
	var form dto.MaterialForm
	err := c.Bind(&form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}

	row, err := srv.AddMaterial(form)
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

func MaterialList(c *gin.Context) {
	var form dto.MaterialSearchForm
	err := c.Bind(&form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	count, list, err := srv.GetMaterialByPage(&form)
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
			"pageIndex": form.PageIndex,
			"PageSize":  form.PageSize,
		},
		"msg": "ok",
	})
	return
}

func GetMaterial(c *gin.Context) {
	Id := c.Param("id")
	mId, _ := strconv.Atoi(Id)
	row, err := srv.GetMaterialById(mId)
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

func EditMaterial(c *gin.Context) {
	var form dto.MaterialForm
	err := c.Bind(&form)
	if err != nil {
		logger.Logger.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	row, err := srv.GetMaterialById(form.Id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}

	row, err = srv.UpdateMaterial(row, form)
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": row,
		"msg":  "ok",
	})
}
