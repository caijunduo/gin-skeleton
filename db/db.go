package db

import (
	"github.com/go-redis/redis"
	"skeleton/config"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
	"upper.io/db.v3/sqlite"
)

func init() {
	if config.DB.Mode {
		if err := NewMySQL(&MySQL.Builder, MySQL.ConnectionURL()); err != nil {
			panic(err)
		}
		if err := NewRedis(&Redis.Builder, Redis.ConnectionURL()); err != nil {
			panic(err)
		}
	}
}

func NewMySQL(builder *sqlbuilder.Database, settings mysql.ConnectionURL) error {
	database, err := mysql.Open(settings)
	if err != nil {
		return err
	}
	database.LoggingEnabled()
	*builder = database
	return nil
}

func NewSQLite(builder *sqlbuilder.Database, settings sqlite.ConnectionURL) error {
	database, err := sqlite.Open(settings)
	if err != nil {
		return err
	}
	*builder = database
	return nil
}

func NewRedis(builder **redis.Client, opt *redis.Options) error {
	engine := redis.NewClient(opt)
	if _, err := engine.Ping().Result(); err != nil {
		return err
	}
	*builder = engine
	return nil
}
