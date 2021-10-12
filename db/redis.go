package db

import (
	"github.com/go-redis/redis"
	"github.com/spf13/cast"
	"skeleton/config"
)

var Redis = redisDB{}

type redisDB struct {
	Builder *redis.Client
}

func (d redisDB) ConnectionURL() *redis.Options {
	return &redis.Options{
		Addr:       config.Redis.Host + ":" + cast.ToString(config.Redis.Port),
		Password:   config.Redis.Auth,
		DB:         config.Redis.Db,
		MaxRetries: config.Redis.MaxRetries,
	}
}
