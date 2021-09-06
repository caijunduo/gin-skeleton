package config

var App = app{}

func init() {
	Include(&App)
}

type app struct {
	Host         string `env:"APP_HOST" envDefault:"127.0.0.1"`
	Port         int    `env:"APP_PORT" envDefault:"3030"`
	ReadTimeout  int    `env:"APP_READ_TIMEOUT" envDefault:"20"`
	WriteTimeout int    `env:"APP_WRITE_TIMEOUT" envDefault:"20"`

	Signature struct {
		AppKey    string `env:"APP_SIGNATURE_APPKEY" envDefault:""`
		AppSecret string `env:"APP_SIGNATURE_APPSECRET" envDefault:""`
		Expires   int    `env:"APP_SIGNATURE_EXPIRES" envDefault:"0"`
	}

	Jwt struct {
		Key           string `env:"APP_JWT_KEY" envDefault:""`
		Issuer        string `env:"APP_JWT_ISSUER" envDefault:""`
		Subject       string `env:"APP_JWT_SUBJECT" envDefault:""`
		ExpireMinutes int    `env:"APP_JWT_EXPIREMINUTES" envDefault:"60"`
	}
}
