package mysqlx

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gohouse/gorose/v2"
)

var (
    engines = make(map[string]*gorose.Engin, 0)
)

type Option struct {
    Connection           string
    Host                 string
    Port                 int
    Username             string
    Password             string
    Database             string
    Prefix               string
    More                 string
    SetMaxIdleConnection int
    SetMaxOpenConnection int
}

func New(opts ...Option) error {
    var (
        err error
    )

    for _, opt := range opts {
        engines[opt.Connection], err = gorose.Open(&gorose.Config{
            Driver: "mysql",
            Dsn: fmt.Sprintf("%s:%s@tcp(%s:%d)/%s%s",
                opt.Username,
                opt.Password,
                opt.Host,
                opt.Port,
                opt.Database,
                opt.More,
            ),
            SetMaxOpenConns: opt.SetMaxOpenConnection,
            SetMaxIdleConns: opt.SetMaxIdleConnection,
            Prefix:          opt.Prefix,
        })
        if err != nil {
            return err
        }
    }

    return nil
}

func Connection(conn string) gorose.IOrm {
    return engines[conn].NewOrm()
}
