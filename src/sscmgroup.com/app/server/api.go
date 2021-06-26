package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"sscmgroup.com/app/config"
	"sscmgroup.com/app/logger"
	"sscmgroup.com/app/server/router"
	"sscmgroup.com/app/service/api"
	"time"
)

var apiServer *http.Server

func InitApi(ctx context.Context, config *config.Config, as *api.Service) (fn func()) {
	apiConf := config.Http["api"]
	engine := gin.Default()
	switch apiConf.Mod {
	case gin.ReleaseMode, gin.TestMode, gin.DebugMode:
		gin.SetMode(apiConf.Mod)
	default:
	}

	apiServer = &http.Server{
		Addr:         apiConf.Addr,
		Handler:      engine,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	router.InitApiRouter(engine, as)
	go func() {
		logger.Logger.Debug("api HTTP server is running at %s.", apiConf.Addr)
		err := apiServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Logger.Error(err.Error())
			panic(err)
		}
	}()

	return func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(10))
		defer cancel()
		apiServer.SetKeepAlivesEnabled(false)
		if err := apiServer.Shutdown(ctx); err != nil {
			logger.Logger.Error(err.Error())
		}
	}
}
