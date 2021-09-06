package config

var Redis = redis{}

func init() {
	Include(&Redis)
}

type redis struct {
	Mode    bool `env:"REDIS_MODE" envDefault:"false"`
	Default struct {
		Host       string `env:"REDIS_HOST" envDefault:"127.0.0.1"`
		Port       int    `env:"REDIS_PORT" envDefault:"6379"`
		Auth       string `env:"REDIS_AUTH" envDefault:""`
		Db         int    `env:"REDIS_DB" envDefault:"0"`
		MaxRetries int    `env:"REDIS_MAXRETRIES" envDefault:"1"`
	}
}
