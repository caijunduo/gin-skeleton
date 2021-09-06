package main

import (
	"github.com/go-redis/redis"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/cast"
	"net/http"
	"skeleton/config"
	"skeleton/pkg/db"
	"skeleton/pkg/errgroup"
	"skeleton/pkg/helper"
	"skeleton/pkg/logger"
	redisPkg "skeleton/pkg/redis"
	"skeleton/pkg/validator"
	"skeleton/routes"
	"time"
	"upper.io/db.v3/mysql"
)

func main() {
	helper.New()
	config.Init()
	logger.New()
	if config.DB.Mode {
		if err := db.NewMySQL(&db.Default, mysql.ConnectionURL{
			User:     config.DB.Default.User,
			Password: config.DB.Default.Password,
			Database: config.DB.Default.Database,
			Host:     config.DB.Default.Host,
			Socket:   config.DB.Default.Socket,
			Options: map[string]string{
				"charset":   config.DB.Default.Options.Charset,
				"parseTime": config.DB.Default.Options.ParseTime,
			},
		}); err != nil {
			panic(err)
		}
	}
	if config.Redis.Mode {
		if err := redisPkg.New(redisPkg.Default, &redis.Options{
			Addr:       config.Redis.Default.Host + ":" + cast.ToString(config.Redis.Default.Port),
			Password:   config.Redis.Default.Auth,
			DB:         config.Redis.Default.Db,
			MaxRetries: config.Redis.Default.MaxRetries,
		}); err != nil {
			panic(err)
		}
	}
	if err := validator.New(); err != nil {
		panic(err)
	}
	s := &http.Server{
		Addr:         config.App.Host + ":" + cast.ToString(config.App.Port),
		Handler:      routes.Init(),
		ReadTimeout:  time.Duration(config.App.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.App.WriteTimeout) * time.Second,
	}
	errgroup.Group.Go(func() error {
		return s.ListenAndServe()
	})
	if err := errgroup.Group.Wait(); err != nil {
		panic(err)
	}
}
