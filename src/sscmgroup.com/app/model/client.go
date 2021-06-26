package model

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"sscmgroup.com/app/config"
	"sscmgroup.com/app/logger"
	"time"
)

var databases map[string]*gorm.DB
var redisClients map[string]*redis.Client

func DbClient(name string) *gorm.DB {
	return databases[name]
}

func RedisClient(name string) *redis.Client {
	return redisClients[name]
}

func InitClient() {
	databases = make(map[string]*gorm.DB)
	redisClients = make(map[string]*redis.Client)
	for key, value := range config.Conf.Databases {
		databases[key] = LoadDb(value)
	}
	for key, value := range config.Conf.Redis {
		redisClients[key] = redis.NewClient(&redis.Options{
			Addr:     value.Addr,
			Password: value.Password,
			PoolSize: value.PoolSize,
		})
	}
}
func LoadDb(c config.Database) *gorm.DB {
	if db, err := gorm.Open(c.Driver, c.Dsn); err != nil {
		panic(err.Error())
	} else {
		db.DB().SetConnMaxLifetime(time.Second * time.Duration(c.IdleTimeout))
		db.DB().SetMaxOpenConns(c.Active)
		db.DB().SetMaxIdleConns(c.Idle)
		db.SetLogger(logger.Logger)
		if c.Debug {
			db.LogMode(true)
			db.Debug()
		}
		return db
	}
}
