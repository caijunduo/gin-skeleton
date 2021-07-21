package main

import (
    _ "github.com/joho/godotenv/autoload"
    _ "github.com/spf13/cast"
    "github.com/spf13/viper"
    "skeleton/internal"
    "skeleton/pkg/configx"
    "skeleton/pkg/helper"
    "skeleton/pkg/logger"
    "skeleton/pkg/mysqlx"
    "skeleton/pkg/redisx"
    "skeleton/pkg/server"
    "skeleton/pkg/validatorx"
)

func main() {
    helper.New()

    if err := configx.New(); err != nil {
        panic(err)
    }

    if viper.GetBool("app.logger") {
        if err := logger.New(); err != nil {
            panic(err)
        }
    }

    if viper.GetBool("app.database") {
        if err := mysqlx.New(); err != nil {
            panic(err)
        }
    }

    if viper.GetBool("app.redis") {
        if err := redisx.New(); err != nil {
            panic(err)
        }
    }

    if err := validatorx.New(); err != nil {
        panic(err)
    }

    server.New(internal.Routes())
}
