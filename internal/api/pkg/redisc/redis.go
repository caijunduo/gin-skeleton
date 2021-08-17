package redisc

import (
    "github.com/go-redis/redis"
    "skeleton/pkg/redisx"
)

func Default() *redis.Client {
    return redisx.Connection("default")
}
