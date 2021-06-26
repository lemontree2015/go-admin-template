package task

import (
	"context"
	"fmt"
	"os"
	"sscmgroup.com/app/config"
	"sscmgroup.com/app/logger"
	"sscmgroup.com/app/model"
)

type options struct {
	ConfigFile string
	TaskName   string
}

// Option 定义配置项
type Option func(*options)

// SetConfigFile 设定配置文件
func SetConfigFile(s string) Option {
	return func(o *options) {
		o.ConfigFile = s
	}
}

// Init 应用初始化
func Init(ctx context.Context, opts ...Option) (fn func(), err error) {
	var o options
	for _, opt := range opts {
		opt(&o)
	}
	config.MustLoad(o.ConfigFile)
	config.PrintWithJSON()
	fmt.Printf("执行Task：%s，进程号：%d", o.TaskName, os.Getpid())
	cleanLogger, err := logger.InitLogger()
	model.InitClient()
	return func() {
		cleanLogger()
	}, err
}
