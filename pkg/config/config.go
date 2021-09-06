package config

import "github.com/caarlos0/env/v6"

func Parse(v interface{}) error {
    if err := env.Parse(v); err != nil {
        return err
    }
    return nil
}
