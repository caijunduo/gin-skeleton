package config

import (
    "github.com/caarlos0/env/v6"
)

func include(opt interface{}) {
    if err := env.Parse(opt); err != nil {
        panic(err)
    }
}
