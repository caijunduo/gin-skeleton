package redisx

import (
    "github.com/go-redis/redis"
)

var Store *redis.Client

func New(opt redis.Options) (*redis.Client, error) {
    Store = redis.NewClient(&opt)
    _, err := Store.Ping().Result()
    if err != nil {
        return Store, err
    }
    return Store, nil
}