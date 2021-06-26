package sys

import (
	"github.com/gin-gonic/gin"
	sysmiddware "sscmgroup.com/app/middleware/sys"
	"sscmgroup.com/app/server/handler/sys"
)

func RegisterSysUserController(r *gin.RouterGroup) {
	profile := r.Group("/profile", sysmiddware.AuthorizedMiddle)
	{
		profile.GET("", sys.GetProfile)
		//用户授权菜单
		profile.GET("/accessSideBar", sys.AccessSideBar)
		profile.PUT("/pwd", sys.ChangePwd)
		profile.POST("/avatar", sys.UploadAvatar)
	}
	user := r.Group("/user", sysmiddware.AuthorizedMiddle, sysmiddware.CasbinMiddleware)
	{
		user.POST("/add", sys.AddUser)
		user.GET("/get/:id", sys.GetUser)
		user.PUT("/edit/:id", sys.EditUser)
		user.DELETE("/delete/:id", sys.DelUser)
		user.GET("/list", sys.UserList)
	}
}
