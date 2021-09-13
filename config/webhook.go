package config

var WebHook = webhook{}

func init() {
	Include(&WebHook)
}

type webhook struct {
	WeCom struct {
		Url string `env:"WEBHOOK_WECOM_URL"`
	}
}
