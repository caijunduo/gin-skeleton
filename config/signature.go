package config

func init() {
	include(&Signature)
}

var Signature struct {
	AppKey    string `env:"APP_SIGNATURE_APPKEY" envDefault:""`
	AppSecret string `env:"APP_SIGNATURE_APPSECRET" envDefault:""`
	Expires   int    `env:"APP_SIGNATURE_EXPIRES" envDefault:"0"`
}
