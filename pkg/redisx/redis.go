package redisx

import (
    "github.com/go-redis/redis"
    "github.com/spf13/cast"
)

var (
    engines = make(map[string]*redis.Client, 0)
)

type Option struct {
    Connection string
    Host       string
    Port       int
    Auth       string
    Db         int
    MaxRetries int
}

func New(opts ...Option) error {
    for _, opt := range opts {
        engines[opt.Connection] = redis.NewClient(&redis.Options{
            Addr:       opt.Host + ":" + cast.ToString(opt.Port),
            Password:   opt.Auth,
            DB:         opt.Db,
            MaxRetries: opt.MaxRetries,
        })
        if _, err := engines[opt.Connection].Ping().Result(); err != nil {
            return err
        }
    }

    return nil
}

func Connection(conn string) *redis.Client {
    return engines[conn]
}
