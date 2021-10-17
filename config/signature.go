package config

func init() {
	include(&Signature)
}

var Signature struct {
	AppKey    string `env:"SIGNATURE_APPKEY" envDefault:""`
	AppSecret string `env:"SIGNATURE_APPSECRET" envDefault:""`
	Expires   int    `env:"SIGNATURE_EXPIRES" envDefault:"0"`
}
