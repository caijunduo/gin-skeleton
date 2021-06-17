package redisx

import (
    "fmt"
    "github.com/go-redis/redis"
    "github.com/spf13/viper"
)

var (
    Store   *redis.Client
    engines = make(map[string]*redis.Client, 0)
)

func New() error {
    var (
        key string
        err error
    )

    r := viper.GetStringMap("redis")

    for conn := range r {
        key = "redis." + conn + "."
        engines[conn] = redis.NewClient(&redis.Options{
            Addr: fmt.Sprintf("%s:%s",
                viper.GetString(key+"host"),
                viper.GetString(key+"port"),
            ),
            Password:   viper.GetString(key + "auth"),
            DB:         viper.GetInt(key + "db"),
            MaxRetries: viper.GetInt(key + "maxRetries"),
        })
        if _, err = engines[conn].Ping().Result(); err != nil {
            return err
        }
    }

    return nil
}

func Connection(conn string) *redis.Client {
    return engines[conn]
}

func Default() *redis.Client {
    return engines["default"]
}
