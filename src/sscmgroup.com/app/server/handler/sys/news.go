package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"net/http"
	"sscmgroup.com/app/config"
	dto "sscmgroup.com/app/dto/sys"
	"sscmgroup.com/app/logger"
	"strconv"
)

func CategoryList(c *gin.Context) {
	var form dto.CategorySearchForm
	err := c.Bind(&form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	count, list, err := srv.GetCategoryByPage(&form)
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

func CategoryAll(c *gin.Context) {
	list, err := srv.GetAllCategories()
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
	return
}

func AddCategory(c *gin.Context) {
	var form dto.NewsCategoryForm
	err := c.Bind(&form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}

	row, err := srv.AddCategory(form)
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

func GetCategory(c *gin.Context) {
	Id := c.Param("id")
	cateId, _ := strconv.Atoi(Id)
	category, err := srv.GetCategoryById(cateId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": category,
		"msg":  "ok",
	})
}

func EditCategory(c *gin.Context) {
	var form dto.NewsCategoryForm
	err := c.Bind(&form)
	if err != nil {
		logger.Logger.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	row, err := srv.GetCategoryById(form.CateId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}

	row, err = srv.UpdateCategory(row, form)
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": row,
		"msg":  "ok",
	})
}

func AddNews(c *gin.Context) {
	var form dto.NewsForm
	err := c.Bind(&form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}

	row, err := srv.AddNews(form)
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

func NewsList(c *gin.Context) {
	var form dto.NewsSearchForm
	err := c.Bind(&form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	count, list, err := srv.GetNewsByPage(&form)
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

func GetNews(c *gin.Context) {
	Id := c.Param("id")
	newsId, _ := strconv.Atoi(Id)
	row, err := srv.GetNewsById(newsId)
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

func EditNews(c *gin.Context) {
	var form dto.NewsForm
	err := c.Bind(&form)
	if err != nil {
		logger.Logger.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	row, err := srv.GetNewsById(form.Id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}

	row, err = srv.UpdateNews(row, form)
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": row,
		"msg":  "ok",
	})
}

func UploadImg(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	uploadCfg := config.Conf.Uploads["news"]

	guid := uuid.New().String()
	fileName := guid + ".jpg"
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
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"url":  uploadCfg.CdnHost + fileName,
		"msg":  "ok",
	})
}
