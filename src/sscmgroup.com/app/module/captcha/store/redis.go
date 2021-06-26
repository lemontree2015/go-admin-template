package store

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/mojocn/base64Captcha"
	"sscmgroup.com/app/logger"
	"time"
)

// NewRedisStore create an instance of a redis store
func NewRedisStore(opts *redis.Options, expiration time.Duration, prefix ...string) base64Captcha.Store {
	if opts == nil {
		panic("options cannot be nil")
	}
	return NewRedisStoreWithCli(
		redis.NewClient(opts),
		expiration,
		prefix...,
	)
}

// NewRedisStoreWithCli create an instance of a redis store
func NewRedisStoreWithCli(cli *redis.Client, expiration time.Duration, prefix ...string) base64Captcha.Store {
	store := &redisStore{
		cli:        cli,
		expiration: expiration,
	}
	if len(prefix) > 0 {
		store.prefix = prefix[0]
	}
	return store
}

// NewRedisClusterStore create an instance of a redis cluster store
func NewRedisClusterStore(opts *redis.ClusterOptions, expiration time.Duration, prefix ...string) base64Captcha.Store {
	if opts == nil {
		panic("options cannot be nil")
	}
	return NewRedisClusterStoreWithCli(
		redis.NewClusterClient(opts),
		expiration,
		prefix...,
	)
}

// NewRedisClusterStoreWithCli create an instance of a redis cluster store
func NewRedisClusterStoreWithCli(cli *redis.ClusterClient, expiration time.Duration, prefix ...string) base64Captcha.Store {
	store := &redisStore{
		cli:        cli,
		expiration: expiration,
	}
	if len(prefix) > 0 {
		store.prefix = prefix[0]
	}
	return store
}

type clienter interface {
	Get(key string) *redis.StringCmd
	Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Del(keys ...string) *redis.IntCmd
}

type redisStore struct {
	cli        clienter
	prefix     string
	expiration time.Duration
}

func (s *redisStore) getKey(id string) string {
	return s.prefix + id
}

func (s *redisStore) printf(format string, args ...interface{}) {
	if logger.Logger != nil {
		logger.Logger.Printf(format, args...)
	}
}

func (s *redisStore) Set(id string, value string) {
	cmd := s.cli.Set(s.getKey(id), value, s.expiration)
	if err := cmd.Err(); err != nil {
		s.printf("redis execution set command error: %s", err.Error())
	}
	return
}

func (s *redisStore) Get(id string, clear bool) string {
	key := s.getKey(id)
	cmd := s.cli.Get(key)
	if err := cmd.Err(); err != nil {
		if err == redis.Nil {
			return ""
		}
		s.printf("redis execution get command error: %s", err.Error())
		return ""
	}
	v := cmd.Val()
	//b, err := hex.DecodeString(cmd.Val())
	//if err != nil {
	//	s.printf("hex decoding error: %s", err.Error())
	//	return ""
	//}

	if clear {
		cmd := s.cli.Del(key)
		if err := cmd.Err(); err != nil {
			s.printf("redis execution del command error: %s", err.Error())
			return ""
		}
	}

	return v
}

//Verify captcha's answer directly
func (s *redisStore) Verify(id, answer string, clear bool) bool {
	cacheCode := s.Get(id, clear)
	fmt.Println("redis Verify", id, cacheCode, answer)
	return cacheCode == answer
}
