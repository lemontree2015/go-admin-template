package logger

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"sscmgroup.com/app/config"
)

var Logger *logrus.Logger

var loggerLevel map[string]logrus.Level

func init() {
	loggerLevel = make(map[string]logrus.Level)
	loggerLevel["panic"] = logrus.PanicLevel
	loggerLevel["fatal"] = logrus.FatalLevel
	loggerLevel["error"] = logrus.ErrorLevel
	loggerLevel["warn"] = logrus.WarnLevel
	loggerLevel["info"] = logrus.InfoLevel
	loggerLevel["debug"] = logrus.DebugLevel
	loggerLevel["trace"] = logrus.TraceLevel
}

func InitLogger() (fn func(), err error) {
	cfg := config.Conf.Log
	Logger = logrus.New()
	//设置输出样式，自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	Logger.SetFormatter(&logrus.JSONFormatter{})
	Logger.SetOutput(os.Stdout)
	//设置output,默认为stderr,可以为任何io.Writer，比如文件*os.File
	if cfg.Dir == "" || cfg.FileName == "" {
		fmt.Println("日志路径未配制")
		return nil, errors.New("日志路径未配制")
	}
	//writers = append(writers, os.Stdout)
	path := cfg.Dir + "/" + cfg.FileName
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	writers := []io.Writer{file, os.Stdout}
	//同时写文件和屏幕
	fileAndStdoutWriter := io.MultiWriter(writers...)
	if err == nil {
		Logger.SetOutput(fileAndStdoutWriter)
	} else {
		Logger.Info("failed to log to file.")
	}
	//设置最低loglevel
	if level, ok := loggerLevel[cfg.Level]; ok {
		Logger.SetLevel(level)
	} else {
		Logger.SetLevel(logrus.InfoLevel)
	}

	fn = func() {
		file.Close()
	}
	return fn, err
}
