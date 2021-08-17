package configs

import (
    "fmt"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "os"
    "time"
)

type configure struct {
    Cmd struct {
        Api struct {
            Host         string        `yaml:"host"`
            Port         int           `yaml:"port"`
            ReadTimeout  time.Duration `yaml:"readTimeout"`
            WriteTimeout time.Duration `yaml:"writeTimeout"`
            Locale       string        `yaml:"locale"`
            Signature    struct {
                AppKey    string `yaml:"appKey"`
                AppSecret string `yaml:"appSecret"`
                Expires   int    `yaml:"expires"`
            } `yaml:"signature"`
            Jwt struct {
                Key           string `yaml:"key"`
                Issuer        string `yaml:"issuer"`
                Subject       string `yaml:"subject"`
                ExpireMinutes int    `yaml:"expireMinutes"`
            } `yaml:"jwt"`
        } `yaml:"api"`
    } `yaml:"cmd"`
    Logger struct {
        Mode       bool   `yaml:"mode"`
        SavePath   string `yaml:"savePath"`
        MaxSize    int    `yaml:"maxSize"`
        MaxAge     int    `yaml:"maxAge"`
        MaxBackups int    `yaml:"maxBackups"`
        Compress   bool   `yaml:"compress"`
    } `yaml:"logger"`
    Cors struct {
        AllowOrigin        string `yaml:"allowOrigin"`
        AllowMethods       string `yaml:"allowMethods"`
        AllowHeaders       string `yaml:"allowHeaders"`
        AllowExposeHeaders string `yaml:"allowExposeHeaders"`
        AllowCredentials   string `yaml:"allowCredentials"`
    } `yaml:"cors"`
    Database struct {
        Mode    bool `yaml:"mode"`
        Default struct {
            Driver               string `yaml:"driver"`
            Host                 string `yaml:"host"`
            Port                 int    `yaml:"port"`
            Username             string `yaml:"username"`
            Password             string `yaml:"password"`
            Database             string `yaml:"database"`
            Prefix               string `yaml:"prefix"`
            More                 string `yaml:"more"`
            SetMaxIdleConnection int    `yaml:"setMaxIdleConnection"`
            SetMaxOpenConnection int    `yaml:"setMaxOpenConnection"`
        } `yaml:"default"`
    } `yaml:"database"`
    Redis struct {
        Mode    bool `yaml:"mode"`
        Default struct {
            Host       string `yaml:"host"`
            Port       int    `yaml:"port"`
            Auth       string `yaml:"auth"`
            Db         int    `yaml:"db"`
            MaxRetries int    `yaml:"maxRetries"`
        } `yaml:"default"`
    } `yaml:"redis"`
    Webhook struct {
        WechatWork struct {
            Url string `yaml:"url"`
        } `yaml:"wechat_work"`
    } `yaml:"webhook"`
}

var Config = &configure{}

func New() {
    c, err := ioutil.ReadFile("config.yaml")
    if err != nil {
        panic("get yaml configure error: " + err.Error())
    }
    c = []byte(os.ExpandEnv(string(c)))
    if err = yaml.Unmarshal(c, Config); err != nil {
        panic(fmt.Sprintf("configure init Unmarshal: %v", err))
    }
}
