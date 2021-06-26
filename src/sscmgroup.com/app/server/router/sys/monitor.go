package sys

import (
	"github.com/gin-gonic/gin"
	sysmiddware "sscmgroup.com/app/middleware/sys"
	"sscmgroup.com/app/server/handler/sys"
)

func RegisterSysMonitorController(r *gin.RouterGroup) {
	menu := r.Group("/monitor", sysmiddware.AuthorizedMiddle)
	{
		menu.GET("/server", sys.ServerInfo)
	}
}
