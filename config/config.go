package config

import (
	"github.com/caarlos0/env/v6"
)

type option interface{}

var options []option

func include(opts ...option) {
	options = append(options, opts...)
}

func Setup() {
	for _, opt := range options {
		if err := env.Parse(opt); err != nil {
			panic(err)
		}
	}
}
