package api

import (
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
)

func CategoryList(c *gin.Context) {
	lang := c.Param("lang")
	parentId := c.Param("pid")
	pId, _ := strconv.Atoi(parentId)
	rows, err := srv.GetSubCateByPid(lang, pId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": rows,
		"msg":  "ok",
	})
}

func CategoryNews(c *gin.Context) {
	cid := c.Param("cateId")
	cpage := c.Param("page")
	page := 1
	cateId, _ := strconv.Atoi(cid)
	if cpage != "" {
		page, _ = strconv.Atoi(cpage)
	}
	rows, count, err := srv.GetCateNews(cateId, page)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code":      http.StatusOK,
		"data":      rows,
		"totalPage": math.Ceil(float64(count) / 10),
		"msg":       "ok",
	})
}

func NewsDetail(c *gin.Context) {
	id := c.Param("id")
	newsId, _ := strconv.Atoi(id)
	row, err := srv.GetNewsDetail(newsId)
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

func NewsAboutUs(c *gin.Context) {
	lang := c.Param("lang")
	cateId := c.Param("cid")
	cId, _ := strconv.Atoi(cateId)
	rows, err := srv.GetAboutUsNews(lang, cId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": rows,
		"msg":  "ok",
	})
}

func AdList(c *gin.Context) {
	position := c.Param("position")
	pId, _ := strconv.Atoi(position)
	rows, err := srv.GetAdListByPosition(pId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": rows,
		"msg":  "ok",
	})
}
