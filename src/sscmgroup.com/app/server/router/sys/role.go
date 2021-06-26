package sys

import (
	"github.com/gin-gonic/gin"
	sysmiddware "sscmgroup.com/app/middleware/sys"
	"sscmgroup.com/app/server/handler/sys"
)

func RegisterSysRoleController(r *gin.RouterGroup) {
	role := r.Group("/role", sysmiddware.AuthorizedMiddle, sysmiddware.CasbinMiddleware)
	{
		role.POST("/add", sys.AddRole)
		role.GET("/get/:id", sys.GetRole)
		role.POST("/edit/:id", sys.EditRole)
		role.GET("/list", sys.RoleList)
		//所有角色列表 用户分配角色使用
		role.GET("/all", sys.AllSysRoles)
		role.GET("/roleMenus/:roleId", sys.GetRoleMenus)
	}
}
