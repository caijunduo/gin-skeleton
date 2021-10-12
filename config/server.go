package config

func init() {
	include(&Server)
}

var Server struct {
	Host         string `env:"SERVER_HOST" envDefault:"0.0.0.0"`
	Port         int    `env:"SERVER_PORT" envDefault:"3030"`
	ReadTimeout  int    `env:"SERVER_READ_TIMEOUT" envDefault:"20"`
	WriteTimeout int    `env:"SERVER_WRITE_TIMEOUT" envDefault:"20"`
}
