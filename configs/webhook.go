package configs

var WebHook = webhook{}

type webhook struct {
    weCom struct {
        Url string `env:"WEBHOOK_WECOM_URL"`
    }
}
