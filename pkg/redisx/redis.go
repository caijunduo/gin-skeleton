package redisx

import (
    "fmt"
    "github.com/go-redis/redis"
    "github.com/spf13/cast"
    "github.com/spf13/viper"
)

var (
    engines = make(map[string]*redis.Client, 0)
)

func New() error {
    var (
        maps map[string]interface{}
        err  error
    )

    r := viper.GetStringMap("redis")
    for connection, v := range r {
        maps = cast.ToStringMap(v)
        engines[connection] = redis.NewClient(&redis.Options{
            Addr: fmt.Sprintf("%s:%s",
                cast.ToString(maps["host"]),
                cast.ToString(maps["port"]),
            ),
            Password:   cast.ToString(maps["auth"]),
            DB:         cast.ToInt(maps["db"]),
            MaxRetries: cast.ToInt(maps["maxRetries"]),
        })
        if _, err = engines[connection].Ping().Result(); err != nil {
            return err
        }
    }

    return nil
}

func Connection(conn string) *redis.Client {
    return engines[conn]
}
