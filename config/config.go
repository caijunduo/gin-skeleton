package config

import configPkg "skeleton/pkg/config"

type option interface{}

var options []option

func Include(opts ...option) {
	options = append(options, opts...)
}

func Init() {
	for _, opt := range options {
		if err := configPkg.Parse(opt); err != nil {
			panic(err)
		}
	}
}
