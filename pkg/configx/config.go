package configx

import (
    "github.com/gin-gonic/gin"
    "github.com/spf13/viper"
    "os"
    "strings"
)

func New() error {
    var (
        key string
        pos int
        val string
    )

    viper.SetConfigName("config")
    viper.AddConfigPath(".")
    viper.SetConfigType("yaml")
    if err := viper.ReadInConfig(); err != nil {
        return err
    }

    for _, k := range viper.AllKeys() {
        value := viper.GetString(k)
        if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
            key = strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}")
            if pos = strings.Index(key, ":"); pos >= 0 {
                if val = os.Getenv(key[:pos]); val == "" {
                    viper.Set(k, key[pos+1:])
                } else {
                    viper.Set(k, val)
                }
            } else {
                viper.Set(k, os.Getenv(key))
            }
        } else {
            viper.Set(k, value)
        }
    }

    if gin.IsDebugging() {
        viper.WatchConfig()
    }

    return nil
}
