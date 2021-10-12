package config

func init() {
	include(&Authorization)
}

var Authorization struct {
	Key           string `env:"Authorization_KEY" envDefault:""`
	Issuer        string `env:"Authorization_ISSUER" envDefault:""`
	Subject       string `env:"Authorization_SUBJECT" envDefault:""`
	ExpireMinutes int    `env:"Authorization_EXPIREMINUTES" envDefault:"60"`
}
