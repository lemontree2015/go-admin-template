package api

import (
	"github.com/gin-gonic/gin"
	"sscmgroup.com/app/server/handler/api"
)

func RegisterNewsController(r *gin.RouterGroup) {
	ad := r.Group("/ad")
	{
		ad.GET("/list/:position", api.AdList)
	}
	news := r.Group("/news")
	{
		news.GET("/list/:cateId/:page", api.CategoryNews)
		news.GET("/detail/:id", api.NewsDetail)
		news.GET("/about/:lang/:cid", api.NewsAboutUs)
	}
	category := r.Group("/category")
	{
		category.GET("/list/:lang/:pid", api.CategoryList)
	}
}
