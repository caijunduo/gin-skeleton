package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/cast"
	"golang.org/x/sync/errgroup"
	"net/http"
	"skeleton/config"
	"skeleton/db"
	_ "skeleton/listener"
	"skeleton/logger"
	"skeleton/router"
	"time"
)

var g errgroup.Group

func main() {
	config.Setup()
	logger.Setup()
	db.Setup()
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
