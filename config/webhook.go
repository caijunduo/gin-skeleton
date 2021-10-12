package config

func init() {
	include(&WebHook)
}

var WebHook struct {
	WeComUrl string `env:"WEBHOOK_WECOMURL"`
}
