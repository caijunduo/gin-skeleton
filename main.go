package main

import (
	_ "github.com/joho/godotenv/autoload" // 自动加载.env环境变量
	"github.com/spf13/cast"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"skeleton/code"
	"skeleton/config"
	"skeleton/router"
	"time"
)

var g errgroup.Group

func main() {
	log.Println(code.OK)

	s := &http.Server{
		Addr:         config.Server.Host + ":" + cast.ToString(config.Server.Port),
		Handler:      router.Setup(),
		ReadTimeout:  time.Duration(config.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.Server.WriteTimeout) * time.Second,
	}
	g.Go(func() error {
		return s.ListenAndServe()
	})
	if err := g.Wait(); err != nil {
		panic(err)
	}
}
