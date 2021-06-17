package mysqlx

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gohouse/gorose/v2"
    "github.com/spf13/viper"
)

var (
    engines = make(map[string]*gorose.Engin, 0)
)

func New() error {
    var (
        k   string
        err error
    )

    r := viper.GetStringMap("database")

    for conn := range r {
        k = "database." + conn + "."
        engines[conn], err = gorose.Open(&gorose.Config{
            Driver: viper.GetString(k + "driver"),
            Dsn: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
                viper.GetString(k+"username"),
                viper.GetString(k+"password"),
                viper.GetString(k+"host"),
                viper.GetString(k+"port"),
                viper.GetString(k+"database"),
                viper.GetString(k+"charset"),
                viper.GetString(k+"parseTime"),
                viper.GetString(k+"loc"),
            ),
            SetMaxOpenConns: viper.GetInt(k + "setMaxOpenConns"),
            SetMaxIdleConns: viper.GetInt(k + "setMaxIdleConns"),
            Prefix:          viper.GetString(k + "prefix"),
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

func Default() gorose.IOrm {
    return engines["default"].NewOrm()
}
