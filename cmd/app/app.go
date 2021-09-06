package main

import (
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"net/http"
	"skeleton/configs"
	appInternal "skeleton/internal/app"
	appDb "skeleton/internal/pkg/db"
	appRedis "skeleton/internal/pkg/redis"
	"skeleton/pkg/config"
	"skeleton/pkg/db"
	"skeleton/pkg/errgroupx"
	"skeleton/pkg/helper"
	"skeleton/pkg/logger"
	redisx "skeleton/pkg/redis"
	"skeleton/pkg/validatorx"
	"time"
	"upper.io/db.v3/mysql"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	helper.New()
	if err := config.Parse(&configs.App); err != nil {
		panic(err)
	}
	if err := config.Parse(&configs.DB); err != nil {
		panic(err)
	}
	if err := config.Parse(&configs.Cors); err != nil {
		panic(err)
	}
	if err := config.Parse(&configs.WebHook); err != nil {
		panic(err)
	}
	logger.New()
	if configs.DB.Mode {
		if err := db.NewMySQL(&appDb.Default, mysql.ConnectionURL{
			User:     configs.DB.Default.User,
			Password: configs.DB.Default.Password,
			Database: configs.DB.Default.Database,
			Host:     configs.DB.Default.Host,
			Socket:   configs.DB.Default.Socket,
			Options: map[string]string{
				"charset":   configs.DB.Default.Options.Charset,
				"parseTime": configs.DB.Default.Options.ParseTime,
			},
		}); err != nil {
			panic(err)
		}
	}
	if configs.Redis.Mode {
		if err := redisx.New(appRedis.Default, &redis.Options{
			Addr:       configs.Redis.Default.Host + ":" + cast.ToString(configs.Redis.Default.Port),
			Password:   configs.Redis.Default.Auth,
			DB:         configs.Redis.Default.Db,
			MaxRetries: configs.Redis.Default.MaxRetries,
		}); err != nil {
			panic(err)
		}
	}
	if err := validatorx.New(); err != nil {
		panic(err)
	}
	s := &http.Server{
		Addr:         configs.App.Host + ":" + cast.ToString(configs.App.Port),
		Handler:      appInternal.Routes(),
		ReadTimeout:  time.Duration(configs.App.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(configs.App.WriteTimeout) * time.Second,
	}
	errgroupx.Group.Go(func() error {
		return s.ListenAndServe()
	})
	if err := errgroupx.Group.Wait(); err != nil {
		zap.L().Error("[Server]", zap.Error(err))
	}
}
