package sys

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sscmgroup.com/app/model"
)

type Repository struct {
	mgrDb     *gorm.DB
	cmsDb     *gorm.DB
	userCache *redis.Client
	bookCache *redis.Client
}

func (r *Repository) UserCache() *redis.Client {
	return r.userCache
}
func (r *Repository) BookCache() *redis.Client {
	return r.bookCache
}
func (r *Repository) MgrDb() *gorm.DB {
	return r.mgrDb
}
func (r *Repository) ApiDb() *gorm.DB {
	return r.cmsDb
}
func New() *Repository {
	rep := &Repository{}
	for _, v := range []string{"mgr_db", "cms_db"} {
		if db := model.DbClient(v); db == nil {
			panic(fmt.Sprintf("%s database config is not found", v))
		} else {
			switch v {
			case "mgr_db":
				rep.mgrDb = db
			case "cms_db":
				rep.cmsDb = db
			}
		}
	}
	for _, v := range []string{"user_cache", "book_cache"} {
		if cache := model.RedisClient(v); cache == nil {
			panic(fmt.Sprintf("%s cache config is not found", v))
		} else {
			switch v {
			case "user_cache":
				rep.userCache = cache
			case "book_cache":
				rep.bookCache = cache
			}
		}
	}
	return rep
}
