package server

import (
    "golang.org/x/sync/errgroup"
    "log"
    "net/http"
    "time"
)

var Group errgroup.Group

func New(port string, handler http.Handler) {
    s := &http.Server{
        Addr:         ":" + port,
        Handler:      handler,
        ReadTimeout:  20 * time.Minute,
        WriteTimeout: 20 * time.Minute,
    }

    Group.Go(func() error {
        return s.ListenAndServe()
    })

    if err := Group.Wait(); err != nil {
        log.Fatal(err)
    }
}
