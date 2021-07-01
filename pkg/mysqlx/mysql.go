package mysqlx

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gohouse/gorose/v2"
    "github.com/spf13/cast"
    "github.com/spf13/viper"
)

var (
    engines = make(map[string]*gorose.Engin, 0)
)

func New() error {
    var (
        maps map[string]interface{}
        err  error
    )

    r := viper.GetStringMap("database")
    for connection, v := range r {
        maps = cast.ToStringMap(v)
        engines[connection], err = gorose.Open(&gorose.Config{
            Driver: cast.ToString(maps["driver"]),
            Dsn: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s%s",
                cast.ToString(maps["username"]),
                cast.ToString(maps["password"]),
                cast.ToString(maps["host"]),
                cast.ToString(maps["port"]),
                cast.ToString(maps["database"]),
                cast.ToString(maps["more"]),
            ),
            SetMaxOpenConns: cast.ToInt(maps["setMaxOpenConns"]),
            SetMaxIdleConns: cast.ToInt(maps["setMaxIdleConns"]),
            Prefix:          cast.ToString(maps["prefix"]),
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
