package router

import (
	"github.com/gin-gonic/gin"
	"sscmgroup.com/app/middleware"
	apihandler "sscmgroup.com/app/server/handler/api"
	syshandler "sscmgroup.com/app/server/handler/sys"
	"sscmgroup.com/app/server/router/api"
	"sscmgroup.com/app/server/router/sys"
	srvapi "sscmgroup.com/app/service/api"
	srvsys "sscmgroup.com/app/service/sys"
)

func InitMgrRouter(eg *gin.Engine, ms *srvsys.MgrService) {
	registerSysRouter(eg, ms)
}

func InitApiRouter(eg *gin.Engine, as *srvapi.Service) {
	registerApiRouter(eg, as)
}

func registerSysRouter(eg *gin.Engine, s *srvsys.MgrService) {
	middleware.CommonMiddleware(eg)
	sysRouter := eg.Group("/")
	syshandler.Init(s)
	sys.RegisterLoginController(sysRouter)
	sys.RegisterSysUserController(sysRouter)
	sys.RegisterSysRoleController(sysRouter)
	sys.RegisterSysMenuController(sysRouter)
	sys.RegisterSysSettingController(sysRouter)
	sys.RegisterNewsController(sysRouter)
	sys.RegisterSysAdMaterialController(sysRouter)
	sys.RegisterSysMonitorController(sysRouter)
}

func registerApiRouter(eg *gin.Engine, s *srvapi.Service) {
	middleware.CommonMiddleware(eg)
	//apiRouter := eg.Group("/api")
	apiRouter := eg.Group("/")
	apihandler.Init(s)
	api.RegisterNewsController(apiRouter)
}
