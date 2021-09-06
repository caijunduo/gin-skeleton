package config

var Cors = cors{}

func init() {
    Include(&Cors)
}

type cors struct {
    AllowOrigin        string `env:"CORS_ALLOWORIGIN" envDefault:"*"`
    AllowMethods       string `env:"CORS_ALLOWMETHODS" envDefault:"POST, GET, OPTIONS, PUT, DELETE, UPDATE"`
    AllowHeaders       string `env:"CORS_ALLOWHEADERS" envDefault:"Origin, X-Requested-With, Content-Type, Accept, Authorization"`
    AllowExposeHeaders string `env:"CORS_ALLOWEXPOSEHEADERS" envDefault:"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type"`
    AllowCredentials   string `env:"CORS_ALLOWCREDENTIALS" envDefault:"true"`
}
