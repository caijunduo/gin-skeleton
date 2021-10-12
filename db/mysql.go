package db

import (
	"skeleton/config"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
)

var MySQL = mysqlDB{}

type mysqlDB struct {
	Builder sqlbuilder.Database
}

func (d mysqlDB) ConnectionURL() mysql.ConnectionURL {
	return mysql.ConnectionURL{
		User:     config.MySQL.User,
		Password: config.MySQL.Password,
		Database: config.MySQL.Database,
		Host:     config.MySQL.Host,
		Socket:   config.MySQL.Socket,
		Options: map[string]string{
			"charset":   config.MySQL.Options.Charset,
			"parseTime": config.MySQL.Options.ParseTime,
		},
	}
}
