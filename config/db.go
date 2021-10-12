package config

func init() {
	include(&MySQL)
	include(&Redis)
	include(&DB)
}

var MySQL struct {
	User     string `env:"MYSQL_USER" envDefault:"root"`
	Password string `env:"MYSQL_PASSWORD" envDefault:""`
	Database string `env:"MYSQL_DATABASE" envDefault:""`
	Host     string `env:"MYSQL_HOST" envDefault:"127.0.0.1:3306"`
	Socket   string `env:"MYSQL_SOCKET" envDefault:""`
	Options  struct {
		Charset   string `env:"MYSQL_OPTIONS_CHARSET" envDefault:"utf8mb4"`
		ParseTime string `env:"MYSQL_OPTIONS_PARSE_TIME" envDefault:"false"`
	}
}

var Redis struct {
	Host       string `env:"REDIS_HOST" envDefault:"127.0.0.1"`
	Port       int    `env:"REDIS_PORT" envDefault:"6379"`
	Auth       string `env:"REDIS_AUTH" envDefault:""`
	Db         int    `env:"REDIS_DB" envDefault:"0"`
	MaxRetries int    `env:"REDIS_MAXRETRIES" envDefault:"1"`
}

var DB struct {
	Mode bool `env:"DB_MODE" envDefault:"false"`
}
