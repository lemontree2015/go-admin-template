package app

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"sscmgroup.com/app/casbin"
	"sscmgroup.com/app/config"
	"sscmgroup.com/app/logger"
	"sscmgroup.com/app/model"
	"sscmgroup.com/app/module/adapter"
	"sscmgroup.com/app/module/captcha"
	"sscmgroup.com/app/server"
	"sscmgroup.com/app/service/sys"
	"syscall"
	"time"
)

// Init 应用初始化
func InitMgr(ctx context.Context, opts ...Option) (fn func(), err error) {
	var o options
	for _, opt := range opts {
		opt(&o)
	}
	config.MustLoad(o.ConfigFile)
	//if v := o.ModelFile; v != "" {
	//	config.Conf.Casbin.Model = v
	//}
	config.PrintWithJSON()
	// 初始化日志模块
	cleanLogger, err := logger.InitLogger()
	// 初始化图形验证码
	model.InitClient()
	captcha.InitCaptcha()
	mgrSrv := sys.New()
	casbinAdapter := &adapter.CasbinAdapter{
		MgrDb: model.DbClient("mgr_db"),
	}
	cleanCasbin, err := casbin.InitCasbin(casbinAdapter)
	cleanHttpServer := server.InitMgr(ctx, config.Conf, mgrSrv)
	logger.Logger.WithFields(logrus.Fields{"v": o.Version, "p": os.Getpid()}).Info("mgr Http server启动")
	return func() {
		cleanLogger()
		cleanCasbin()
		cleanHttpServer()
	}, err
}

// Run 运行服务
func RunMgr(ctx context.Context, opts ...Option) error {
	state := 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	cleanFunc, err := InitMgr(ctx, opts...)
	if err != nil {
		return err
	}

EXIT:
	for {
		sig := <-sc
		//fmt.Println("接收到信号[%s]", sig.String())
		logger.Logger.WithContext(ctx).Infof("接收到信号[%s]", sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			state = 0
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}

	logger.Logger.WithContext(ctx).Infof("服务退出")
	//fmt.Println("服务退出")
	cleanFunc()
	time.Sleep(time.Second)
	os.Exit(state)
	return nil
}
