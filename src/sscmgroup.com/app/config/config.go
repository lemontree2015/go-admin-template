package config

import (
	"encoding/json"
	"github.com/koding/multiconfig"
	"os"
	"strings"
	"sync"
)

const HttpModRelease string = "release"

var (
	Conf = new(Config)
	once sync.Once
)

type Config struct {
	PrintConfig bool
	Http        map[string]HttpServerConfig
	Log         LogConfig
	Databases   map[string]Database
	Casbin      Casbin
	Jwt         map[string]JwtConfig
	Captcha     Captcha
	Redis       map[string]RedisClientConf
	Uploads     map[string]Upload
}

type Upload struct {
	Dir string
	CdnHost string
}
type HttpServerConfig struct {
	Addr string
	Mod  string
}
type LogConfig struct {
	Dir      string
	FileName string
	Level    string
}
type Database struct {
	Driver      string
	Dsn         string
	Active      int
	Idle        int
	IdleTimeout int
	Debug       bool
}

// Casbin casbin配置参数
type Casbin struct {
	Enable           bool
	Debug            bool
	Model            string
	AutoLoad         bool
	AutoLoadInternal int
}

type JwtConfig struct {
	Secret string
	TTL    int `toml:"ttl"`
}

// Captcha 图形验证码配置参数
type Captcha struct {
	Store       string
	Length      int
	Width       int
	Height      int
	RedisDB     int
	RedisPrefix string
	TTL         int
}

type RedisClientConf struct {
	Addr     string `toml:"addrs"`
	Password string
	PoolSize int `toml:"pool_size"`
}

// MustLoad 加载配置
func MustLoad(fpaths ...string) {
	once.Do(func() {
		loaders := []multiconfig.Loader{
			&multiconfig.TagLoader{},
			&multiconfig.EnvironmentLoader{},
		}

		for _, fpath := range fpaths {
			if strings.HasSuffix(fpath, "toml") {
				loaders = append(loaders, &multiconfig.TOMLLoader{Path: fpath})
			}
			if strings.HasSuffix(fpath, "json") {
				loaders = append(loaders, &multiconfig.JSONLoader{Path: fpath})
			}
			if strings.HasSuffix(fpath, "yaml") {
				loaders = append(loaders, &multiconfig.YAMLLoader{Path: fpath})
			}
		}

		m := multiconfig.DefaultLoader{
			Loader:    multiconfig.MultiLoader(loaders...),
			Validator: multiconfig.MultiValidator(&multiconfig.RequiredValidator{}),
		}
		m.MustLoad(Conf)
	})
}

// PrintWithJSON 基于JSON格式输出配置
func PrintWithJSON() {
	if Conf.PrintConfig {
		b, err := json.MarshalIndent(Conf, "", " ")
		if err != nil {
			os.Stdout.WriteString("[CONFIG] JSON marshal error: " + err.Error())
			return
		}
		os.Stdout.WriteString(string(b) + "\n")
	}
}
