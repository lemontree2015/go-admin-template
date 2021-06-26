package sys

import (
	"github.com/gin-gonic/gin"
	sysmiddware "sscmgroup.com/app/middleware/sys"
	"sscmgroup.com/app/server/handler/sys"
)

func RegisterSysMenuController(r *gin.RouterGroup) {
	menu := r.Group("/menu", sysmiddware.AuthorizedMiddle, sysmiddware.CasbinMiddleware)
	{
		menu.POST("/add", sys.AddMenu)
		menu.GET("/list", sys.MenusList)
		menu.GET("/get/:id", sys.GetMenu)
		menu.PUT("/edit/:id", sys.EditMenu)
	}
}
