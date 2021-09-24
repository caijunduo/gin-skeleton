package cache

import (
	"github.com/go-redis/redis"
	"skeleton/pkg"
	"skeleton/response"
	"time"
)

func Remember(key string, ttl time.Duration, callback func() (interface{}, response.Response)) (interface{}, response.Response) {
	var res interface{}
	if err := pkg.RedisDefault.Get(key).Scan(&res); err == redis.Nil {
		if resSet, errSet := callback(); errSet != nil {
			return nil, errSet
		} else {
			pkg.RedisDefault.Set(key, resSet, ttl)
			return resSet, nil
		}
	} else if err != nil {
		return nil, response.InternalServerError.SetError(err)
	} else {
		return res, nil
	}
}
