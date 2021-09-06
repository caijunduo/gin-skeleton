package config

var DB = db{}

func init() {
    Include(&DB)
}

type db struct {
    Mode    bool `env:"DB_MODE" envDefault:"false"`
    Default struct {
        User     string `env:"DB_MYSQL_USER" envDefault:"root"`
        Password string `env:"DB_MYSQL_PASSWORD" envDefault:""`
        Database string `env:"DB_MYSQL_DATABASE" envDefault:""`
        Host     string `env:"DB_MYSQL_HOST" envDefault:"127.0.0.1:3306"`
        Socket   string `env:"DB_MYSQL_SOCKET" envDefault:""`
        Options  struct {
            Charset   string `env:"DB_MYSQL_OPTIONS_CHARSET" envDefault:"utf8mb4"`
            ParseTime string `env:"DB_MYSQL_OPTIONS_PARSE_TIME" envDefault:"false"`
        }
    }
}
