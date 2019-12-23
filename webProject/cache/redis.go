package cache

import (
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

type redisH struct {
	client *redis.Client
}

func (r redisH) Set(name string, value string, options interface{}) error {
	if expirationTime, ok := options.(time.Duration); !ok {
		panic("参数错误")
	} else {
		status := r.client.Set(name, value, expirationTime)
		if _, err := status.Result(); err != nil {
			return err
		}
		return nil
	}
}

func (r redisH) Get(name string) (string, error) {

	status := r.client.Get(name)
	if val, err := status.Result(); err != nil {
		return val, nil
	} else {
		return val, err
	}
}

func InitRedis(db string, addr string, pwd string) *redisH {
	var r = new(redisH)
	d, _ := strconv.ParseUint(db, 10, 64)
	r.client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       int(d),
	})

	return r
}
