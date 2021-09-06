package redis

import (
    "github.com/go-redis/redis"
)

func New(builder *redis.Client, opt *redis.Options) error {
    engine := redis.NewClient(opt)
    if _, err := engine.Ping().Result(); err != nil {
        return err
    }
    builder = engine
    return nil
}
