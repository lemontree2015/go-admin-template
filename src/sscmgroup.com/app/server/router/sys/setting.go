package sys

import (
	"github.com/gin-gonic/gin"
	sysmiddware "sscmgroup.com/app/middleware/sys"
	"sscmgroup.com/app/server/handler/sys"
)

func RegisterSysSettingController(r *gin.RouterGroup) {
	menu := r.Group("/setting", sysmiddware.AuthorizedMiddle, sysmiddware.CasbinMiddleware)
	{
		menu.POST("/add", sys.AddSetting)
		menu.GET("/list", sys.SettingList)
		menu.GET("/get/:id", sys.GetSetting)
		menu.PUT("/edit/:id", sys.EditSetting)
	}
}
