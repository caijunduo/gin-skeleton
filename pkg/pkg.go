package pkg

import (
	"github.com/go-redis/redis"
	"golang.org/x/sync/errgroup"
	"upper.io/db.v3/lib/sqlbuilder"
)

var (
	Group        errgroup.Group
	RedisDefault *redis.Client
	MySQLDefault sqlbuilder.Database
)
