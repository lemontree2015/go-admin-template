package sys

import (
	"github.com/gin-gonic/gin"
	sysmiddware "sscmgroup.com/app/middleware/sys"
	"sscmgroup.com/app/server/handler/sys"
)

func RegisterNewsController(r *gin.RouterGroup) {
	news := r.Group("/news", sysmiddware.AuthorizedMiddle, sysmiddware.CasbinMiddleware)
	category := news.Group("/category")
	{
		category.GET("/list", sys.CategoryList)
		category.GET("/all", sys.CategoryAll)
		category.POST("/add", sys.AddCategory)
		category.GET("/get/:id", sys.GetCategory)
		category.PUT("/edit/:id", sys.EditCategory)
	}
	publish := news.Group("/publish")
	{
		publish.POST("/add", sys.AddNews)
		publish.GET("/list", sys.NewsList)
		publish.GET("/get/:id", sys.GetNews)
		publish.PUT("/edit/:id", sys.EditNews)
		publish.POST("/uploadImg", sys.UploadImg)
	}
}
