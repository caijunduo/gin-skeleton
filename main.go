package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/cast"
	"log"
	"net/http"
	"skeleton/config"
	"skeleton/internal"
	"skeleton/pkg"
	_ "skeleton/pkg"
	"skeleton/pkg/db"
	"skeleton/pkg/logger"
	"skeleton/pkg/structure"
	"time"
	"upper.io/db.v3/mysql"
)

func main() {
	var ss structure.StringSlice = []string{"1", "", "2"}
	log.Println(ss.FilterNilAndCreate())
	log.Println(ss)
	log.Println(ss.FilterNil().Output())
	log.Println(ss.Is("2"))
	config.Init()
	logger.New()
	if config.DB.Mode {
		if err := db.NewMySQL(&pkg.MySQLDefault, mysql.ConnectionURL{
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
		if err := db.NewRedis(&pkg.RedisDefault, &redis.Options{
			Addr:       config.Redis.Default.Host + ":" + cast.ToString(config.Redis.Default.Port),
			Password:   config.Redis.Default.Auth,
			DB:         config.Redis.Default.Db,
			MaxRetries: config.Redis.Default.MaxRetries,
		}); err != nil {
			panic(err)
		}
	}
	s := &http.Server{
		Addr: config.App.Host + ":" + cast.ToString(config.App.Port),
		Handler: func() *gin.Engine {
			r := gin.New()
			root := internal.Root{}
			root.RouteEngine(r)
			root.RouteGroup(r.Group(""))
			return r
		}(),
		ReadTimeout:  time.Duration(config.App.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.App.WriteTimeout) * time.Second,
	}
	pkg.Group.Go(func() error {
		return s.ListenAndServe()
	})
	if err := pkg.Group.Wait(); err != nil {
		panic(err)
	}
}
