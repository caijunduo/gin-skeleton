package main

import (
    "github.com/joho/godotenv"
    "github.com/spf13/cast"
    "go.uber.org/zap"
    "net/http"
    "skeleton/configs"
    apiInternal "skeleton/internal/api"
    "skeleton/pkg/errgroupx"
    "skeleton/pkg/helper"
    "skeleton/pkg/logger"
    "skeleton/pkg/mysqlx"
    "skeleton/pkg/redisx"
    "skeleton/pkg/validatorx"
    "time"
)

func main() {
    if err := godotenv.Load(); err != nil {
        panic(err)
    }
    helper.New()
    configs.New()
    logger.New()
    if configs.Config.Database.Mode {
        if err := mysqlx.New(mysqlx.Option{
            Connection:           "default",
            Host:                 configs.Config.Database.Default.Host,
            Port:                 configs.Config.Database.Default.Port,
            Username:             configs.Config.Database.Default.Username,
            Password:             configs.Config.Database.Default.Password,
            Database:             configs.Config.Database.Default.Database,
            Prefix:               configs.Config.Database.Default.Prefix,
            More:                 configs.Config.Database.Default.More,
            SetMaxIdleConnection: configs.Config.Database.Default.SetMaxIdleConnection,
            SetMaxOpenConnection: configs.Config.Database.Default.SetMaxOpenConnection,
        }); err != nil {
            panic(err)
        }
    }
    if configs.Config.Redis.Mode {
        if err := redisx.New(redisx.Option{
            Connection: "default",
            Host:       configs.Config.Redis.Default.Host,
            Port:       configs.Config.Redis.Default.Port,
            Auth:       configs.Config.Redis.Default.Auth,
            Db:         configs.Config.Redis.Default.Db,
            MaxRetries: configs.Config.Redis.Default.MaxRetries,
        }); err != nil {
            panic(err)
        }
    }
    if err := validatorx.New(configs.Config.Cmd.Api.Locale); err != nil {
        panic(err)
    }

    s := &http.Server{
        Addr:         configs.Config.Cmd.Api.Host + ":" + cast.ToString(configs.Config.Cmd.Api.Port),
        Handler:      apiInternal.Routes(),
        ReadTimeout:  configs.Config.Cmd.Api.ReadTimeout * time.Minute,
        WriteTimeout: configs.Config.Cmd.Api.WriteTimeout * time.Minute,
    }

    errgroupx.Group.Go(func() error {
        return s.ListenAndServe()
    })

    if err := errgroupx.Group.Wait(); err != nil {
        zap.L().Error("[Server]", zap.Error(err))
    }
}
