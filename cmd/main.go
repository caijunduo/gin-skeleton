package main

import (
    "fmt"
    "github.com/go-redis/redis"
    _ "github.com/joho/godotenv/autoload"
    "github.com/spf13/cast"
    _ "github.com/spf13/cast"
    "os"
    "skeleton/internal"
    "skeleton/pkg/logger"
    "skeleton/pkg/mysql"
    "skeleton/pkg/redisx"
    "skeleton/pkg/server"
)

func main() {
    if cast.ToBool(os.Getenv("LOG_SWITCH")) {
        if err := logger.New(); err != nil {
            panic(err)
        }
    }

    if cast.ToBool(os.Getenv("DB_SWITCH")) {
        if _, err := mysql.New(mysql.Dsn{
            Driver:   "mysql",
            Host:     os.Getenv("DB_HOST"),
            Port:     os.Getenv("DB_PORT"),
            Username: os.Getenv("DB_USERNAME"),
            Password: os.Getenv("DB_PASSWORD"),
            Database: os.Getenv("DB_DATABASE"),
            Charset:  os.Getenv("DB_CHARSET"),
        }); err != nil {
            panic(err)
        }
    }

    if cast.ToBool(os.Getenv("REDIS_SWITCH")) {
        if _, err := redisx.New(redis.Options{
            Addr:       fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
            Password:   os.Getenv("REDIS_AUTH"),
            DB:         cast.ToInt(os.Getenv("REDIS_DB")),
            MaxRetries: 1,
        }); err != nil {
            panic(err)
        }
    }

    if cast.ToBool(os.Getenv("CRONTAB_SWITCH")) {
        server.Group.Go(func() error {
            return internal.Crontab()
        })
    }

    server.New(os.Getenv("APP_PORT"), internal.Routes())
}
