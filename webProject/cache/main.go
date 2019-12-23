package cache

import (
	"os"
)

type cacheHandle interface {
	Set(name string, value string, options interface{}) error
	Get(name string) (string, error)
}

var Handle cacheHandle

func NewCacheHandle(s string) {

	switch s {
	case "Redis":
		c := InitRedis(os.Getenv("REDIS_DB"), os.Getenv("REDIS_ADDR"), os.Getenv("REDIS_PW"))
		Handle = c
	}

}
