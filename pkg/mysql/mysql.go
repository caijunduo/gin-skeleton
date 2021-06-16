package mysql

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gohouse/gorose/v2"
)

var (
    engines = make(map[string]*gorose.Engin, 0)
    err     error
)

type Dsn struct {
    Driver   string
    Host     string
    Port     string
    Username string
    Password string
    Database string
    Charset  string
}

func (d Dsn) String() string {
    return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
        d.Username,
        d.Password,
        d.Host,
        d.Port,
        d.Database,
        d.Charset,
    )
}

func New(dsn ...Dsn) (map[string]*gorose.Engin, error) {
    for _, v := range dsn {
        engines[v.Host+v.Database], err = gorose.Open(&gorose.Config{Driver: v.Driver, Dsn: v.String()})
        if err != nil {
            return nil, err
        }
    }
    return engines, nil
}

func DB(host string, database string) gorose.IOrm {
    return engines[host+database].NewOrm()
}
