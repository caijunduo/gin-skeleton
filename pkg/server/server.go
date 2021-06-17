package server

import (
    "github.com/spf13/viper"
    "go.uber.org/zap"
    "golang.org/x/sync/errgroup"
    "net/http"
    "time"
)

var g errgroup.Group

func New(handler http.Handler) {
    s := &http.Server{
        Addr:         ":" + viper.GetString("app.port"),
        Handler:      handler,
        ReadTimeout:  time.Duration(viper.GetInt("app.readTimeout")) * time.Minute,
        WriteTimeout: time.Duration(viper.GetInt("app.writeTimeout")) * time.Minute,
    }

    g.Go(func() error {
        return s.ListenAndServe()
    })

    if err := g.Wait(); err != nil {
        zap.L().Error("[Server]", zap.Error(err))
    }
}

func Group() *errgroup.Group {
    return &g
}
