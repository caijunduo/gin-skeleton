package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/cast"
	"net/http"
	"skeleton/config"
	"skeleton/db"
	"skeleton/pkg"
	_ "skeleton/pkg"
	"skeleton/router"
	"time"
)

func main() {
	config.Setup()
	pkg.Logger()
	db.Setup()
	s := &http.Server{
		Addr:         config.Server.Host + ":" + cast.ToString(config.Server.Port),
		Handler:      router.Setup(),
		ReadTimeout:  time.Duration(config.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.Server.WriteTimeout) * time.Second,
	}
	pkg.Group.Go(func() error {
		return s.ListenAndServe()
	})
	if err := pkg.Group.Wait(); err != nil {
		panic(err)
	}
}
