package config

func init() {
	include(&Authorization)
}

var Authorization struct {
	Key           string `env:"AUTHORIZATION_KEY" envDefault:""`
	Issuer        string `env:"AUTHORIZATION_ISSUER" envDefault:""`
	Subject       string `env:"AUTHORIZATION_SUBJECT" envDefault:""`
	ExpireMinutes int    `env:"AUTHORIZATION_EXPIREMINUTES" envDefault:"60"`
}
