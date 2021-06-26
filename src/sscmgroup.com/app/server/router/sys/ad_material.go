package sys

import (
	"github.com/gin-gonic/gin"
	sysmiddware "sscmgroup.com/app/middleware/sys"
	"sscmgroup.com/app/server/handler/sys"
)

func RegisterSysAdMaterialController(r *gin.RouterGroup) {
	material := r.Group("/material", sysmiddware.AuthorizedMiddle, sysmiddware.CasbinMiddleware)
	{
		material.POST("/add", sys.AddMaterial)
		material.GET("/list", sys.MaterialList)
		material.GET("/get/:id", sys.GetMaterial)
		material.PUT("/edit/:id", sys.EditMaterial)
	}
}
