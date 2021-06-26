package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"sscmgroup.com/app/config"
	"sscmgroup.com/app/logger"
	"sscmgroup.com/app/server/router"
	"sscmgroup.com/app/service/sys"
	"time"
)

var mgrServer *http.Server

func InitMgr(ctx context.Context, config *config.Config, ms *sys.MgrService) (fn func()) {
	mgrConf := config.Http["mgr"]
	engine := gin.Default()
	switch mgrConf.Mod {
	case gin.ReleaseMode, gin.TestMode, gin.DebugMode:
		gin.SetMode(mgrConf.Mod)
	default:
	}

	mgrServer = &http.Server{
		Addr:         mgrConf.Addr,
		Handler:      engine,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	router.InitMgrRouter(engine, ms)
	go func() {
		logger.Logger.Debug("mgr HTTP server is running at %s.", mgrConf.Addr)
		err := mgrServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Logger.Error(err.Error())
			panic(err)
		}
	}()

	return func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(10))
		defer cancel()
		mgrServer.SetKeepAlivesEnabled(false)
		if err := mgrServer.Shutdown(ctx); err != nil {
			logger.Logger.Error(err.Error())
		}
	}
}
